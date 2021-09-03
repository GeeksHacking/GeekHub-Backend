package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/payload"
	"github.com/geekshacking/geekhub-backend/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type ticket struct {
	usecase usecase.Ticket
}

func NewTicket(usecase usecase.Ticket) *ticket {
	return &ticket{usecase}
}

func (t *ticket) NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", t.FindByProjectID)
	r.Post("/", t.Create)
	return r
}

func (t *ticket) FindByProjectID(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParamFromCtx(r.Context(), "ID")

	ID, err := strconv.Atoi(projectID)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid project ID"})
		return
	}

	tickets, err := t.usecase.FindByProjectID(r.Context(), ID)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
	}

	response := make([]payload.TicketResponse, 0, len(tickets))
	for _, ticket := range tickets {
		response = append(response, payload.TicketResponse{
			ID:             ticket.ID,
			Name:           ticket.Name,
			Description:    ticket.Description,
			Type:           ticket.Type.String(),
			Status:         ticket.Status.String(),
			ReporterID:     derefReporter(ticket.Edges).Auth0ID,
			AssigneeID:     derefAssignee(ticket.Edges).Auth0ID,
			ParentTicketID: derefParent(ticket.Edges).ID,
		})
	}

	render.DefaultResponder(w, r, render.M{"data": response})
}

func (t *ticket) Create(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParamFromCtx(r.Context(), "ID")

	ID, err := strconv.Atoi(projectID)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid project ID"})
		return
	}

	var body payload.CreateTicketRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid body"})
		return
	}

	if err = body.Validate(); err != nil {
		render.DefaultResponder(w, r, render.M{"error": "bad request", "details": err.Error()})
		return
	}

	result, err := t.usecase.Create(
		r.Context(),
		ID,
		body.Name,
		body.Description,
		body.Type,
		body.Status,
		body.ReporterID,
		body.AssigneeID,
		body.ParentTicketID,
	)
	if errors.Is(err, usecase.ErrValidation) {
		render.DefaultResponder(w, r, render.M{"error": "bad request"})
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
	}

	render.DefaultResponder(w, r, render.M{"data": payload.TicketResponse{
		ID:             result.ID,
		Name:           result.Name,
		Description:    result.Description,
		Type:           result.Type.String(),
		Status:         result.Status.String(),
		ReporterID:     derefReporter(result.Edges).Auth0ID,
		AssigneeID:     derefAssignee(result.Edges).Auth0ID,
		ParentTicketID: derefParent(result.Edges).ID,
	}})
}

func derefReporter(r ent.TicketEdges) ent.User {
	if r.Reporter == nil {
		var zero ent.User
		return zero
	}
	return *r.Reporter
}

func derefAssignee(r ent.TicketEdges) ent.User {
	if r.Assignee == nil {
		var zero ent.User
		return zero
	}
	return *r.Reporter
}

func derefParent(r ent.TicketEdges) ent.Ticket {
	if r.Parent == nil {
		var zero ent.Ticket
		return zero
	}
	return *r.Parent
}
