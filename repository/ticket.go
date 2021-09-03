package repository

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
)

type TicketReader interface {
	Find(ctx context.Context, ID int) (*ent.Ticket, error)
	FindByProjectID(ctx context.Context, ID int) ([]*ent.Ticket, error)
}

type TicketWriter interface {
	Create(ctx context.Context, ticket ent.Ticket) (*ent.Ticket, error)
	Update(ctx context.Context, ticket ent.Ticket) (*ent.Ticket, error)
}

type Ticket interface {
	TicketReader
	TicketWriter
}