package usecase

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/geekshacking/geekhub-backend/config"
	"github.com/geekshacking/geekhub-backend/ent"
	"github.com/geekshacking/geekhub-backend/repository"

	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Project interface {
	Find(ctx context.Context, ID int) (*ent.Project, error)
	FindByUserAuth0ID(ctx context.Context, userID string) ([]*ent.Project, error)
	Create(
		ctx context.Context,
		name string,
		description string,
		repository string,
		ownerID string,
	) (*ent.Project, error)
	Update(
		ctx context.Context,
		ID int,
		name string,
		description string,
		repository string,
	) (*ent.Project, error)
}

type project struct {
	config             config.Config
	repository         repository.Project
	languageRepository repository.Language
	userRepository     repository.User
	httpClient         http.Client
}

func NewProject(config config.Config, repository repository.Project, languageRepository repository.Language, userRepository repository.User) Project {
	httpClient := http.Client{}
	return &project{config, repository, languageRepository, userRepository, httpClient}
}

func (p *project) Find(ctx context.Context, ID int) (*ent.Project, error) {
	return p.repository.Find(ctx, ID)
}

func (p *project) FindByUserAuth0ID(ctx context.Context, userID string) ([]*ent.Project, error) {
	return p.repository.FindByUserAuth0ID(ctx, userID)
}

func (p *project) Create(
	ctx context.Context,
	name string,
	description string,
	repository string,
	ownerID string,
) (*ent.Project, error) {
	user, err := p.userRepository.FindByAuth0ID(ctx, ownerID)

	var notFoundError *ent.NotFoundError
	if errors.As(err, &notFoundError) {
		name, err := p.getUserDisplayName(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not get user details: %w", err)
		}

		user, err = p.userRepository.Create(ctx, ent.User{
			ID:          0,
			DisplayName: name,
			Auth0ID:     ownerID,
		})
		if err != nil {
			return nil, fmt.Errorf("could not create user: %w", err)
		}
	}

	// Check valid repo and also get languages
	languages, err := p.getRepositoryLanguages(ctx, repository)
	if err != nil {
		return nil, ErrInvalidGitHubRepository
	}

	result, err := p.repository.Create(ctx, ent.Project{
		Name:        name,
		Description: description,
		Repository:  repository,
		Edges: ent.ProjectEdges{
			Owner: user,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("could not create project: %w", err)
	}

	var newLanguages []*ent.Language
	var existingLanguages []*ent.Language
	for _, name := range languages {
		// Language does not exist
		var notFoundError *ent.NotFoundError
		language, err := p.languageRepository.FindByName(ctx, name)
		if errors.As(err, &notFoundError) {
			newLanguages = append(newLanguages, &ent.Language{
				Name: name,
				Edges: ent.LanguageEdges{
					Projects: []*ent.Project{result},
				},
			})
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("could not find language: %w", err)
		}

		existingLanguages = append(existingLanguages, language)
	}

	_, err = p.languageRepository.CreateBulk(ctx, newLanguages)
	if err != nil {
		return nil, fmt.Errorf("could not create languages: %w", err)
	}

	_, err = p.repository.BulkAddLanguage(ctx, result.ID, existingLanguages)
	if err != nil {
		return nil, fmt.Errorf("could not create languages: %w", err)
	}

	return result, nil
}

func (p *project) Update(
	ctx context.Context,
	ID int,
	name string,
	description string,
	repository string,
) (*ent.Project, error) {
	languages, err := p.getRepositoryLanguages(ctx, repository)
	if err != nil {
		return nil, ErrInvalidGitHubRepository
	}

	entLanguages := make([]*ent.Language, 0, len(languages))
	for idx, language := range languages {
		entLanguages[idx] = &ent.Language{
			Name: language,
		}
	}

	entLanguages, err = p.languageRepository.CreateBulk(ctx, entLanguages)
	if err != nil {
		return nil, fmt.Errorf("could not create languages: %w", err)
	}

	result, err := p.repository.Update(ctx, ent.Project{
		ID:          ID,
		Name:        name,
		Description: description,
		Repository:  repository,
		Edges: ent.ProjectEdges{
			Languages: entLanguages,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("could not update project: %w", err)
	}

	return result, nil
}

func (p *project) getRepositoryLanguages(ctx context.Context, repository string) ([]string, error) {
	repositoryPathSplit := strings.Split(repository, "/")
	if len(repositoryPathSplit) < 2 {
		return nil, errors.New("invalid repository link")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", repositoryPathSplit[len(repositoryPathSplit)-2], repositoryPathSplit[len(repositoryPathSplit)-1]), nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %w", err)
	}

	res, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while doing request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != 200 {
		return nil, errors.New("invalid repository link")
	}

	var body map[string]int
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, fmt.Errorf("error while decoding response from github API: %w", err)
	}

	keys := make([]string, 0, len(body))
	for key := range body {
		keys = append(keys, key)
	}

	return keys, nil
}

func (p *project) getUserDisplayName(ctx context.Context) (string, error) {
	token := ctx.Value("user").(*jwt.Token).Raw

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%suserinfo", p.config.Domain), nil)
	if err != nil {
		return "", fmt.Errorf("error while creating request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error while doing request: %w", err)
	}

	var body map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return "", fmt.Errorf("error while decoding response from Auth0 API: %w", err)
	}

	if _, ok := body["name"]; !ok {
		return "Unknown", nil
	}

	return body["name"].(string), nil
}
