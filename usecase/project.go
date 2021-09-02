package usecase

import (
	"context"
	"fmt"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/repository"
)

type Project interface {
	Find(ctx context.Context, ID int) (*ent.Project, error)
	FindByUserID(ctx context.Context, userID string) ([]*ent.Project, error)
	Create(ctx context.Context, model ent.Project) (*ent.Project, error)
	Update(ctx context.Context, model ent.Project) (*ent.Project, error)
}

type project struct {
	repository repository.Project
}

func NewProject(repository repository.Project) *project {
	return &project{repository}
}

func (p *project) Find(ctx context.Context, ID int) (*ent.Project, error) {
	result, err := p.repository.Find(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("could not find project with ID %d: %w", ID, err)
	}

	return result, nil
}

func (p *project) FindByUserID(ctx context.Context, userID string) ([]*ent.Project, error) {
	return p.repository.FindByUserID(ctx, userID)
}

func (p *project) Create(ctx context.Context, model ent.Project) (*ent.Project, error) {
	panic("")
}

func (p *project) Update(ctx context.Context, model ent.Project) (*ent.Project, error) {
	panic("")
}
