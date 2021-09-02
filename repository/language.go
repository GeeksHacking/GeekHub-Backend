package repository

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
)

type LanguageWriter interface {
	CreateBulk(ctx context.Context, models []*ent.Language) ([]*ent.Language, error)
}

type Language interface {
	LanguageWriter
}
