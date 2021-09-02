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
	result, err := p.client.Project.Query().Where(entproject.ID(ID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *project) FindByUserAuth0ID(ctx context.Context, userID string) ([]*ent.Project, error) {
	result, err := p.client.User.Query().Where(entuser.Auth0ID(userID)).QueryProjects().All(ctx)
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
