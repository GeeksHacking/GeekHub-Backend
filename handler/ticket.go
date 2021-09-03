package handler

import (
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
			TicketType:     ticket.Type.String(),
			TicketStatus:  	ticket.Status.String(),
			ReporterID:     ticket.Edges.Reporter.ID,
			AssigneeID:     ticket.Edges.Assignee.ID,
			ParentTicketID: ticket.Edges.Parent.ID,
		})
	}

	render.DefaultResponder(w, r, render.M{"data": response})
}

