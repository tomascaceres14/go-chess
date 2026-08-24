package user

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrUsernameTooShort   = errors.New("Username must be at least 6 characters long")
	ErrPasswordTooShort   = errors.New("Password must be at least 8 characters long")
	ErrPasswordsDontMatch = errors.New("Passwords do not match")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(ctx context.Context, register RegisterUser) (*User, error) {
	if err := register.Validate(); err != nil {
		return nil, err
	}

	if register.Password != register.RepeatPassword {
		return nil, ErrPasswordsDontMatch
	}

	user, err := NewUser(register.Username, register.Password)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("Repository error: %w", err)
	}

	return user, nil
}

func (s *Service) GetAll(ctx context.Context) []*User {
	return s.repo.GetAll(ctx)
}
