package postgres

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
	entproject "github.com/geekshacking/geekhub-backend/ent/project"
	entuser "github.com/geekshacking/geekhub-backend/ent/user"
	"github.com/geekshacking/geekhub-backend/repository"
)

type project struct {
	client *ent.Client
}

func NewProject(client *ent.Client) repository.Project {
	return &project{client}
}

func (p *project) Find(ctx context.Context, ID int) (*ent.Project, error) {
	result, err := p.client.Project.Query().
		Where(entproject.ID(ID)).
		WithTags().
		WithOwner().
		WithUsers().
		WithLanguages().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *project) FindByUserAuth0ID(ctx context.Context, userID string) ([]*ent.Project, error) {
	result, err := p.client.Project.Query().
		Where(entproject.HasUsersWith(entuser.Auth0ID(userID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *project) Create(ctx context.Context, model ent.Project) (*ent.Project, error) {
	result, err := p.client.Project.Create().
		SetName(model.Name).
		SetDescription(model.Description).
		SetRepository(model.Repository).
		SetOwner(model.Edges.Owner).
		AddUsers(model.Edges.Users...).
		AddTags(model.Edges.Tags...).
		AddTickets(model.Edges.Tickets...).
		AddLanguages(model.Edges.Languages...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *project) Update(ctx context.Context, model ent.Project) (*ent.Project, error) {
	result, err := p.client.Project.UpdateOneID(model.ID).
		SetName(model.Name).
		SetDescription(model.Description).
		SetRepository(model.Repository).
		AddUsers(model.Edges.Users...).
		AddTags(model.Edges.Tags...).
		AddTickets(model.Edges.Tickets...).
		AddLanguages(model.Edges.Languages...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *project) BulkAddLanguage(ctx context.Context, ID int, models []*ent.Language) (*ent.Project, error) {
	return p.client.Project.UpdateOneID(ID).AddLanguages(models...).Save(ctx)
}
