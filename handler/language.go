package handler

import (
	"github.com/geekshacking/geekhub-backend/payload"
	"github.com/geekshacking/geekhub-backend/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type language struct {
	usecase usecase.Language
}

func NewLanguage(usecase usecase.Language) *language {
	return &language{usecase}
}

func (l *language) NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", l.FindByProjectID)
	return r
}

func (l *language) FindByProjectID(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParamFromCtx(r.Context(), "ID")

	ID, err := strconv.Atoi(projectID)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid project ID"})
		return
	}

	languages, err := l.usecase.FindByProjectID(r.Context(), ID)
	res := make([]payload.LanguageResponse, 0, len(languages))
	for _, language := range languages {
		res = append(res, payload.LanguageResponse{
			ID:   language.ID,
			Name: language.Name,
		})
	}

	render.DefaultResponder(w, r, render.M{"data": res})
}
