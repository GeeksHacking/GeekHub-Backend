package repository

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
)

type TicketReader interface {
	FindByProjectID(ctx context.Context, ID int) ([]*ent.Ticket, error)
}

type Ticket interface {
	TicketReader
}