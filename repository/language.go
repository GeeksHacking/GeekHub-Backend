package repository

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
)

type LanguageReader interface {
	FindByName(ctx context.Context, name string) (*ent.Language, error)
}

type LanguageWriter interface {
	CreateBulk(ctx context.Context, models []*ent.Language) ([]*ent.Language, error)
	AddProject(ctx context.Context, ID int, project *ent.Project) (*ent.Language, error)
}

type Language interface {
	LanguageReader
	LanguageWriter
}
