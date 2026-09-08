package auth

import (
	"context"

	"github.com/tomascaceres14/go-chess/api/internal/token"
	"github.com/tomascaceres14/go-chess/api/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userSvc       *user.Service
	tokenProvider token.TokenProvider
}

func NewService(userService *user.Service) *Service {
	return &Service{
		userSvc: userService,
	}
}

func (s *Service) Register(ctx context.Context, register UserRegister) (*token.AccessCredentials, error) {

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

	credentials, err := s.tokenProvider.NewAccessCredentials(user.ID)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func (s *Service) Login(ctx context.Context, login UserLogin) (*token.AccessCredentials, error) {

	user, err := s.userSvc.GetByUsername(ctx, login.Username)
	if err != nil {
		return nil, err
	}

	if !CheckPasswordHash(login.Password, user.HashedPassword) {
		return nil, err
	}

	credentials, err := s.tokenProvider.NewAccessCredentials(user.ID)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
