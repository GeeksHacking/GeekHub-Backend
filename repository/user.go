package repository

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
)

type UserReader interface {
	FindByAuth0ID(ctx context.Context, ID string) (*ent.User, error)
}

type UserWriter interface {
	Create(ctx context.Context, model ent.User) (*ent.User, error)
}

type User interface {
	UserReader
	UserWriter
}
