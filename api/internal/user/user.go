package user

import (
	"context"
)

type User struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"-"`
}

type Repository interface {
	Save(ctx context.Context, user *User) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	ExistsByID(ctx context.Context, id string) bool
}

func NewUser(username, hashedPassword string) (*User, error) {
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
	}, nil
}
