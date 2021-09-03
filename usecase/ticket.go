package usecase

import (
	"context"
	"fmt"
	"github.com/geekshacking/geekhub-backend/ent"
	entticket "github.com/geekshacking/geekhub-backend/ent/ticket"
	"github.com/geekshacking/geekhub-backend/repository"
)

type Ticket interface {
	FindByProjectID(ctx context.Context, ID int) ([]*ent.Ticket, error)
	Create(
		ctx context.Context,
		projectID int,
		name string,
		description string,
		ticketType string,
		status string,
		reporterID string,
		assigneeID string,
		parentID int,
	) (*ent.Ticket, error)
}

type ticket struct {
	repository        repository.Ticket
	userRepository    repository.User
	projectRepository repository.Project
}

func NewTicket(repository repository.Ticket, userRepository repository.User, projectRepository repository.Project) Ticket {
	return &ticket{repository, userRepository, projectRepository}
}

func (t *ticket) FindByProjectID(ctx context.Context, ID int) ([]*ent.Ticket, error) {
	return t.repository.FindByProjectID(ctx, ID)
}

func (t *ticket) Create(
	ctx context.Context,
	projectID int,
	name string,
	description string,
	ticketType string,
	status string,
	reporterID string,
	assigneeID string,
	parentID int,
) (*ent.Ticket, error) {
	p, err := t.projectRepository.Find(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not find project: %w", err)
	}

	err = entticket.TypeValidator(entticket.Type(ticketType))
	if err != nil {
		return nil, ErrValidation
	}

	err = entticket.StatusValidator(entticket.Status(status))
	if err != nil {
		return nil, ErrValidation
	}

	parent, err := t.repository.Find(ctx, parentID)
	if parentID != 0 && err != nil {
		return nil, fmt.Errorf("could not find parent ticket: %w", err)
	}

	reporter, err := t.userRepository.FindByAuth0ID(ctx, reporterID)
	if len(reporterID) != 0 && err != nil {
		return nil, fmt.Errorf("could not find reporter: %w", err)
	}

	assignee, err := t.userRepository.FindByAuth0ID(ctx, assigneeID)
	if len(assigneeID) != 0 && err != nil {
		return nil, fmt.Errorf("could not find assignee: %w", err)
	}

	result, err := t.repository.Create(ctx, ent.Ticket{
		Name:        name,
		Description: description,
		Type:        entticket.Type(ticketType),
		Status:      entticket.Status(status),
		Edges: ent.TicketEdges{
			Project:  p,
			Reporter: reporter,
			Assignee: assignee,
			Parent:   parent,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("could not create ticket: %w", err)
	}

	return result, nil
}
