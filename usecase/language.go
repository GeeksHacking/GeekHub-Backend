package usecase

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/repository"
)

type Language interface {
	FindByProjectID(ctx context.Context, projectID int) ([]*ent.Language, error)
}

type language struct {
	repository repository.Language
}

func NewLanguage(repository repository.Language) Language {
	return &language{repository}
}

func (l *language) FindByProjectID(ctx context.Context, projectID int) ([]*ent.Language, error) {
	return l.repository.FindByProjectID(ctx, projectID)
}
