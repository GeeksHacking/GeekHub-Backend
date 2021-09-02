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

func (p *project) FindByUserID(ctx context.Context, userID string) ([]*ent.Project, error) {
	result, err := p.client.User.Query().Where(entuser.Auth0ID(userID)).QueryProjects().All(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *project) Create(
	ctx context.Context,
	name string,
	description string,
	repository string,
	owner *ent.User,
	users []*ent.User,
	tags []*ent.Tag,
	tickets []*ent.Ticket,
	languages []*ent.Language,
) (*ent.Project, error) {
	result, err := p.client.Project.Create().
		SetName(name).
		SetDescription(description).
		SetRepository(repository).
		SetOwner(owner).
		AddUsers(users...).
		AddTags(tags...).
		AddTickets(tickets...).
		AddLanguages(languages...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *project) Update(
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
) (*ent.Project, error) {
	result, err := p.client.Project.UpdateOneID(ID).
		SetName(name).
		SetDescription(description).
		SetRepository(repository).
		SetOwner(owner).
		AddUsers(users...).
		AddTags(tags...).
		AddTickets(tickets...).
		AddLanguages(languages...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
