package repository

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
)

type ProjectReader interface {
	Find(ctx context.Context, ID int) (*ent.Project, error)
	FindByUserID(ctx context.Context, userID string) ([]*ent.Project, error)
}

type ProjectWriter interface {
	Create(
		ctx context.Context,
		name string,
		description string,
		repository string,
		owner *ent.User,
		users []*ent.User,
		tags []*ent.Tag,
		tickets []*ent.Ticket,
		languages []*ent.Language,
	) (*ent.Project, error)
	Update(
		ctx context.Context,
		ID int,
		name string,
		description string,
		repository string,
		owner *ent.User,
		users []*ent.User,
		tags []*ent.Tag,
		tickets []*ent.Ticket,
		languages []*ent.Language,
	) (*ent.Project, error)
}

type Project interface {
	ProjectReader
	ProjectWriter
}
