package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/tomascaceres14/go-chess/api/internal/match"
)

type Service struct {
	repo     Repository
	matchSvc *match.Service
}

func NewService(repo Repository, svc *match.Service) *Service {
	return &Service{
		repo:     repo,
		matchSvc: svc,
	}
}

func (s *Service) CreateUser(ctx context.Context, username, hashedPassword string) (*User, error) {
	user, err := NewUser(username, hashedPassword)
	if err != nil {
		return nil, err
	}

	if user, err = s.repo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("Repository error: %w", err)
	}

	return user, nil
}

func (s *Service) GetAll(ctx context.Context) ([]*User, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (*User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) ExistsByID(ctx context.Context, id string) bool {
	return s.repo.ExistsByID(ctx, id)
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*User, error) {

	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("No users.")
	}

	for _, u := range users {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, ErrUserNotFound
}
