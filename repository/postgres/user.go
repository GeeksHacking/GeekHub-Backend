package postgres

import (
	"context"
	"github.com/geekshacking/geekhub-backend/ent"
	entuser "github.com/geekshacking/geekhub-backend/ent/user"
	"github.com/geekshacking/geekhub-backend/repository"
)

type user struct {
	client *ent.Client
}

func NewUser(client *ent.Client) repository.User {
	return &user{client}
}

func (u *user) FindByAuth0ID(ctx context.Context, ID string) (result *ent.User, err error) {
	result, err = u.client.User.Query().Where(entuser.Auth0ID(ID)).Only(ctx)
	return
}

func (u *user) Create(ctx context.Context, model ent.User) (result *ent.User, err error) {
	result, err = u.client.User.Create().
		SetAuth0ID(model.Auth0ID).
		SetDisplayName(model.DisplayName).
		AddProjects(model.Edges.Projects...).
		AddAssignedTickets(model.Edges.AssignedTickets...).
		AddReportedTickets(model.Edges.ReportedTickets...).
		Save(ctx)
	return
}
