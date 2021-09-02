package postgres

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
	entlanguage "github.com/geekshacking/geekhub-backend/ent/language"
	"github.com/geekshacking/geekhub-backend/repository"
)

type language struct {
	client *ent.Client
}

func NewLanguage(client *ent.Client) repository.Language {
	return &language{client}
}

func (l *language) FindByName(ctx context.Context, name string) (*ent.Language, error) {
	return l.client.Language.Query().Where(entlanguage.Name(name)).Only(ctx)
}

func (l *language) CreateBulk(ctx context.Context, models []*ent.Language) ([]*ent.Language, error) {
	bulk := make([]*ent.LanguageCreate, 0, len(models))
	for _, model := range models {
		bulk = append(bulk, l.client.Language.Create().SetName(model.Name).AddProjects(model.Edges.Projects...))
	}

	result, err := l.client.Language.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (l *language) AddProject(ctx context.Context, ID int, project *ent.Project) (*ent.Language, error) {
	return l.client.Language.UpdateOneID(ID).AddProjects(project).Save(ctx)
}