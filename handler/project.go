package handler

import (
	"bytes"
	jsonpatch "github.com/evanphx/json-patch"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/payload"
	"github.com/geekshacking/geekhub-backend/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"encoding/json"
	"errors"
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
	r.Get("/{ID}/languages", p.FindLanguages)
	r.Get("/user/{userID}", p.FindByUserID)
	r.Post("/", p.Create)
	r.Patch("/{ID}", p.Update)
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

func (p *project) FindLanguages(w http.ResponseWriter, r *http.Request) {
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

	languages, err := result.QueryLanguages().All(r.Context())
	res := make([]payload.LanguageResponse, 0, len(languages))
	for _, language := range languages {
		res = append(res, payload.LanguageResponse{
			ID:   language.ID,
			Name: language.Name,
		})
	}

	render.DefaultResponder(w, r, render.M{"data": res})
}

func (p *project) FindByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	result, err := p.usecase.FindByUserAuth0ID(r.Context(), userID)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
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

func (p *project) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	var body payload.CreateProjectRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid body"})
		return
	}

	if err = body.Validate(); err != nil {
		render.DefaultResponder(w, r, render.M{"error": "bad request", "details": err.Error()})
		return
	}

	result, err := p.usecase.Create(r.Context(), body.Name, body.Description, body.Repository, userID)
	if errors.Is(err, usecase.ErrInvalidGitHubRepository) {
		render.DefaultResponder(w, r, render.M{"error": "invalid repository"})
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

func (p *project) Update(w http.ResponseWriter, r *http.Request) {
	_ = r.Context().Value("userID").(string)
	projectIDParam := chi.URLParamFromCtx(r.Context(), "ID")

	projectID, err := strconv.Atoi(projectIDParam)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid project ID"})
		return
	}

	buffer := &bytes.Buffer{}
	_, err = buffer.ReadFrom(r.Body)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid body"})
		return
	}

	patch, err := jsonpatch.DecodePatch(buffer.Bytes())
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid JSON patch"})
		return
	}

	project, err := p.usecase.Find(r.Context(), projectID)
	var notFoundError *ent.NotFoundError
	if errors.As(err, &notFoundError) {
		render.DefaultResponder(w, r, render.M{"error": "could not find project"})
		return
	}
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
	}

	projectJSON, err := json.Marshal(project)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
	}

	modifiedJSON, err := patch.Apply([]byte(projectJSON))
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "could not apply update"})
		return
	}

	var modified payload.UpdateProjectRequest
	err = json.Unmarshal(modifiedJSON, &modified)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "could not apply update"})
		return
	}

	if err = modified.Validate(); err != nil {
		render.DefaultResponder(w, r, render.M{"error": "bad request", "details": err.Error()})
		return
	}

	result, err := p.usecase.Update(r.Context(), projectID, modified.Name, modified.Description, modified.Repository)
	if errors.Is(err, usecase.ErrInvalidGitHubRepository) {
		render.DefaultResponder(w, r, render.M{"error": "invalid repository"})
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
