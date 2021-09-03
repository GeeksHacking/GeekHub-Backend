package usecase

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/repository"
)

type Ticket interface {
	FindByProjectID(ctx context.Context, ID int) ([]*ent.Ticket, error)
}

type ticket struct {
	repository repository.Ticket
}

func NewTicket(repository repository.Ticket) Ticket {
	return &ticket{repository}
}

func (t *ticket) FindByProjectID(ctx context.Context, ID int) ([]*ent.Ticket, error) {
	return t.repository.FindByProjectID(ctx, ID)
}