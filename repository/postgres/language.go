package postgres

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/repository"
)

type language struct {
	client *ent.Client
}

func NewLanguage(client *ent.Client) repository.Language {
	return &language{client}
}

func (l *language) CreateBulk(ctx context.Context, models []*ent.Language) ([]*ent.Language, error) {
	bulk := make([]*ent.LanguageCreate, len(models))
	for idx, model := range models {
		bulk[idx] = l.client.Language.Create().SetName(model.Name).AddProjects(model.Edges.Projects...)
	}

	result, err := l.client.Language.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}


