package postgres

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
	entproject "github.com/geekshacking/geekhub-backend/ent/project"
	entticket "github.com/geekshacking/geekhub-backend/ent/ticket"
	"github.com/geekshacking/geekhub-backend/repository"
)

type ticket struct {
	client *ent.Client
}

func NewTicket(client *ent.Client) repository.Ticket {
	return &ticket{client}
}

func (t *ticket) FindByProjectID(ctx context.Context, ID int) ([]*ent.Ticket, error) {
	return t.client.Ticket.Query().WithProject().Where(
		entticket.HasProjectWith(
			entproject.ID(ID),
		)).All(ctx)
}
