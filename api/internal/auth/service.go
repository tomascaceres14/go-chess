package auth

import (
	"context"

	"github.com/tomascaceres14/go-chess/api/internal/user"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := HashPassword(register.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.userSvc.CreateUser(ctx, register.Username, hashedPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
