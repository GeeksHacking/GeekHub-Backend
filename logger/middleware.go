package logger

import (
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type middleware struct {
	logger *zap.SugaredLogger
}

func NewMiddleware(logger *zap.SugaredLogger) *middleware {
	return &middleware{
		logger: logger,
	}
}

func (m *middleware) Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, r *http.Request) {
		ww := chimiddleware.NewWrapResponseWriter(res, r.ProtoMajor)

		start := time.Now()
		defer func() {
			end := time.Now()
			go func() {
				m.logger.Infof("HTTP %s %s responded %d with %d bytes in %dms", r.Method, r.URL.String(), ww.Status(), ww.BytesWritten(), end.Sub(start)/1000000)
			}()
		}()

		next.ServeHTTP(ww, r)
	})
}
