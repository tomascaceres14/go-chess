package user

import "context"

type MemoryRepository struct {
	users map[string]*User
}

func (r *MemoryRepository) Save(ctx context.Context, user *User) error {
	return nil
}
