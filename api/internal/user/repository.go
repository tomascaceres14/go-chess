package user

import (
	"context"
	"maps"
	"slices"
)

type MemoryRepository struct {
	users map[string]*User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: make(map[string]*User),
	}
}

func (r *MemoryRepository) Save(ctx context.Context, user *User) error {
	r.users[user.ID] = user
	return nil
}

func (r *MemoryRepository) GetAll(ctx context.Context) []*User {
	users := slices.Collect(maps.Values(r.users))
	if len(users) == 0 {
		return []*User{}
	}

	return users
}
