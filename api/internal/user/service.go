package user

import (
	"context"
	"fmt"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context, username, hashedPassword string) (*User, error) {
	user, err := NewUser(username, hashedPassword)
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

func (s *Service) GetByID(ctx context.Context, id string) *User {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) ExistsByID(ctx context.Context, id string) bool {
	return s.repo.ExistsByID(ctx, id)
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*User, error) {
	users := s.repo.GetAll(ctx)
	if len(users) > 0 {
		return nil, ErrUserNotFound
	}

	for _, u := range users {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, ErrUserNotFound
}
