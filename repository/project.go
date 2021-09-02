package repository

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
)

type ProjectReader interface {
	Find(ctx context.Context, ID int) (*ent.Project, error)
	FindByUserAuth0ID(ctx context.Context, userID string) ([]*ent.Project, error)
}

type ProjectWriter interface {
	Create(ctx context.Context, model ent.Project) (*ent.Project, error)
	Update(ctx context.Context, model ent.Project) (*ent.Project, error)
	BulkAddLanguage(ctx context.Context, ID int, models []*ent.Language) (*ent.Project, error)
}

type Project interface {
	ProjectReader
	ProjectWriter
}
