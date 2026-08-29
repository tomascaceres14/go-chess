package auth

import (
	"context"

	"github.com/tomascaceres14/go-chess/api/internal/user"
)

type Service struct {
	userSvc *user.Service
}

func NewService(userService *user.Service) *Service {
	return &Service{
		userSvc: userService,
	}
}

func (s *Service) Register(ctx context.Context, register UserRegister) (*user.User, error) {

	if err := register.Validate(); err != nil {
		return nil, err
	}

	if register.Password != register.RepeatPassword {
		return nil, ErrPasswordsDontMatch
	}

	user, err := s.userSvc.CreateUser(ctx, register.Username, register.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
