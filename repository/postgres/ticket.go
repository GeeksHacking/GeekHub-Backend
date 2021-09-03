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

func (t *ticket) Find(ctx context.Context, ID int) (*ent.Ticket, error) {
	return t.client.Ticket.Query().
		Where(entticket.ID(ID)).
		WithProject().
		WithAssignee().
		WithReporter().
		WithParent().
		WithChildren().
		Only(ctx)
}

func (t *ticket) FindByProjectID(ctx context.Context, ID int) ([]*ent.Ticket, error) {
	return t.client.Ticket.Query().
		Where(entticket.HasProjectWith(entproject.ID(ID))).
		WithProject().
		WithAssignee().
		WithReporter().
		WithParent().
		WithChildren().
		All(ctx)
}

func (t *ticket) Create(ctx context.Context, ticket ent.Ticket) (*ent.Ticket, error) {
	query := t.client.Ticket.Create().
		SetName(ticket.Name).
		SetDescription(ticket.Description).
		SetType(ticket.Type).
		SetStatus(ticket.Status).
		SetProject(ticket.Edges.Project)

	if ticket.Edges.Reporter != nil {
		query.SetReporter(ticket.Edges.Reporter)
	}

	if ticket.Edges.Assignee != nil {
		query.SetAssignee(ticket.Edges.Assignee)
	}

	if ticket.Edges.Parent != nil {
		query.SetParent(ticket.Edges.Parent)
	}

	return query.Save(ctx)
}

func (t *ticket) Update(ctx context.Context, ticket ent.Ticket) (*ent.Ticket, error) {
	query := t.client.Ticket.UpdateOneID(ticket.ID).
		SetName(ticket.Name).
		SetDescription(ticket.Description).
		SetType(ticket.Type).
		SetStatus(ticket.Status)

	if ticket.Edges.Reporter != nil {
		query.SetReporter(ticket.Edges.Reporter)
	}

	if ticket.Edges.Assignee != nil {
		query.SetAssignee(ticket.Edges.Assignee)
	}

	if ticket.Edges.Parent != nil {
		query.SetParent(ticket.Edges.Parent)
	}

	return query.Save(ctx)
}
