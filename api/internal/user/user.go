package user

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"-"`
}

type Repository interface {
	Save(ctx context.Context, user *User) error
	GetAll(ctx context.Context) []*User
	GetByID(ctx context.Context, id string) *User
	ExistsByID(ctx context.Context, id string) bool
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
