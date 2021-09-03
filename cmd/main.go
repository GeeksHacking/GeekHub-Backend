package main

import (
	"github.com/geekshacking/geekhub-backend/config"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/handler"
	"github.com/geekshacking/geekhub-backend/jwt"
	"github.com/geekshacking/geekhub-backend/logger"
	"github.com/geekshacking/geekhub-backend/repository/postgres"
	"github.com/geekshacking/geekhub-backend/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	serverConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	serverLogger, err := logger.NewLogger(serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	serverLogger.Zap.Info("Connecting to database...")

	// Database
	client, err := ent.Open(serverConfig.DriverName, serverConfig.DataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	serverLogger.Zap.Info("Connected to database")

	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	serverLogger.Zap.Info("Creating schema...")

	// Create schema
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Schema.Create(ctx); err != nil {
		serverLogger.Zap.Fatalf("failed creating schema resources: %v", err)
	}

	serverLogger.Zap.Info("Created schema")

	serverLogger.Zap.Info("Creating router...")

	// Services
	r := chi.NewRouter()
	loggerMiddleware := logger.NewMiddleware(serverLogger.Zap)

	jwtMiddleware, err := jwt.NewMiddleware(serverConfig)
	if err != nil {
		serverLogger.Zap.Fatalf("failed creating JWT middleware: %v", err)
	}

	r.Use(jwtMiddleware.Auth0().Handler)
	r.Use(jwtMiddleware.User)
	r.Use(loggerMiddleware.Request)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	userRepository := postgres.NewUser(client)
	projectRepository := postgres.NewProject(client)
	languageRepository := postgres.NewLanguage(client)
	ticketRepository := postgres.NewTicket(client)

	projectUseCase := usecase.NewProject(serverConfig, projectRepository, languageRepository, userRepository)
	languageUseCase := usecase.NewLanguage(languageRepository)
	ticketUseCase := usecase.NewTicket(ticketRepository)

	projectHandler := handler.NewProject(projectUseCase)
	languageHandler := handler.NewLanguage(languageUseCase)
	ticketHandler := handler.NewTicket(ticketUseCase)

	r.Mount("/projects", projectHandler.NewRouter())
	r.Mount("/projects/{ID}/languages", languageHandler.NewRouter())
	r.Mount("/projects/{ID}/tickets", ticketHandler.NewRouter())

	server := http.Server{Addr: serverConfig.ApplicationUrl, Handler: r}

	serverLogger.Zap.Infof("Created router and listening on %s", serverConfig.ApplicationUrl)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			serverLogger.Zap.Fatalf("Failed trying to listen to address: %v", err)
		}
		wg.Done()
	}()

	go func() {
		<-interrupt

		serverLogger.Zap.Info("Shutting down...")

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = server.Shutdown(ctx)
		if err != nil {
			serverLogger.Zap.Fatalf("Error while shutting down: %v", err)
		}

		err = client.Close()
		if err != nil {
			serverLogger.Zap.Fatalf("Error while closing DB connection: %v", err)
		}
	}()

	wg.Wait()

	serverLogger.Zap.Info("Shut down successfully")
}
