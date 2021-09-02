package handler

import (
	"errors"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/payload"
	"github.com/geekshacking/geekhub-backend/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type project struct {
	usecase usecase.Project
}

func NewProject(usecase usecase.Project) *project {
	return &project{usecase}
}

func (p *project) NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/{ID}", p.Find)
	r.Get("/user/{userID}", p.FindByUserID)
	return r
}

func (p *project) Find(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParamFromCtx(r.Context(), "ID")

	ID, err := strconv.Atoi(projectID)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid project ID"})
		return
	}

	result, err := p.usecase.Find(r.Context(), ID)
	var notFoundError *ent.NotFoundError
	if errors.As(err, &notFoundError) {
		render.DefaultResponder(w, r, render.M{"error": "could not find project"})
		return
	}

	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
	}

	render.DefaultResponder(w, r, render.M{"data": payload.ProjectResponse{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Repository:  result.Repository,
	}})
}

func (p *project) FindByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	result, err := p.usecase.FindByUserID(r.Context(), userID)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
	}

	projects := make([]payload.ProjectResponse, 0, len(result))
	for _, project := range result {
		projects = append(projects, payload.ProjectResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Repository:  project.Repository,
		})
	}

	render.DefaultResponder(w, r, render.M{"data": projects})
}