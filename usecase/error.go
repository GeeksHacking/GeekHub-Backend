package usecase

import "errors"

var (
	ErrInvalidGitHubRepository = errors.New("invalid GitHub repository")
	ErrValidation = errors.New("validation error")
)
