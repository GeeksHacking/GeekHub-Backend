package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch"
	"github.com/geekshacking/geekhub-backend/ent"
	entticket "github.com/geekshacking/geekhub-backend/ent/ticket"
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
	r.Get("/{ticketID}", t.Find)
	r.Get("/types", t.FindTicketTypes)
	r.Get("/statuses", t.FindTicketStatuses)
	r.Post("/", t.Create)
	r.Patch("/{ticketID}", t.Update)
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

func (t *ticket) Find(w http.ResponseWriter, r *http.Request) {
	ticketIDString := chi.URLParamFromCtx(r.Context(), "ticketID")
	ticketID, err := strconv.Atoi(ticketIDString)

	ticket, err := t.usecase.Find(r.Context(), ticketID)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
	}

	render.DefaultResponder(w, r, render.M{"data": payload.TicketResponse{
		ID:             ticket.ID,
		Name:           ticket.Name,
		Description:    ticket.Description,
		Type:           ticket.Type.String(),
		Status:         ticket.Status.String(),
		ReporterID:     derefReporter(ticket.Edges).Auth0ID,
		AssigneeID:     derefAssignee(ticket.Edges).Auth0ID,
		ParentTicketID: derefParent(ticket.Edges).ID,
	}})
}

func (t *ticket) FindTicketTypes(w http.ResponseWriter, r *http.Request) {
	render.DefaultResponder(w, r, render.M{"data": []entticket.Type{entticket.TypeEpic, entticket.TypeBug, entticket.TypeStory, entticket.TypeTask}})
}

func (t *ticket) FindTicketStatuses(w http.ResponseWriter, r *http.Request) {
	render.DefaultResponder(w, r, render.M{"data": []entticket.Status{entticket.StatusBacklog, entticket.StatusDevelopment, entticket.StatusQa, entticket.StatusRelease}})
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
		fmt.Println(err)
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

func (t *ticket) Update(w http.ResponseWriter, r *http.Request) {
	ticketIDParam := chi.URLParamFromCtx(r.Context(), "ticketID")

	ticketID, err := strconv.Atoi(ticketIDParam)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "invalid ticket ID"})
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

	ticket, err := t.usecase.Find(r.Context(), ticketID)
	if ent.IsNotFound(err) {
		render.DefaultResponder(w, r, render.M{"error": "could not find ticket"})
		return
	}
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
	}

	ticketJSON, err := json.Marshal(ticket)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "internal server error"})
		return
	}

	modifiedJSON, err := patch.Apply(ticketJSON)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "could not apply update"})
		return
	}

	var modified payload.UpdateTicketRequest
	err = json.Unmarshal(modifiedJSON, &modified)
	if err != nil {
		render.DefaultResponder(w, r, render.M{"error": "could not apply update"})
		return
	}

	if err = modified.Validate(); err != nil {
		render.DefaultResponder(w, r, render.M{"error": "bad request", "details": err.Error()})
		return
	}

	result, err := t.usecase.Update(r.Context(), ticketID, modified.Name, modified.Description, modified.Type, modified.Status, modified.ReporterID, modified.AssigneeID, modified.ParentTicketID)
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
