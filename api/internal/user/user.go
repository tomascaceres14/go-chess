package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUsernameTooShort   = errors.New("Username must be at least 6 characters long")
	ErrPasswordTooShort   = errors.New("Password must be at least 8 characters long")
	ErrPasswordsDontMatch = errors.New("Passwords do not match")
)

type User struct {
	ID             string
	Username       string
	HashedPassword string
}

type Repository interface {
	Save(ctx context.Context, user *User) error
}

func NewUser(username, hashedPassword string) (*User, error) {

	if username == "" {
		return nil, ErrUsernameTooShort
	}

	return &User{
		ID:             uuid.NewString(),
		Username:       username,
		HashedPassword: hashedPassword,
	}, nil
}
