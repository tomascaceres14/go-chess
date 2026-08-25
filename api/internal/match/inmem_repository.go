package match

import (
	"context"
)

// TESTING PURPOSES ONLY
type MemoryRepository struct {
	Matches map[string]*Match
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		Matches: make(map[string]*Match),
	}
}

func (r *MemoryRepository) Save(ctx context.Context, Match *Match) error {
	r.Matches[Match.ID] = Match
	return nil
}

func (r *MemoryRepository) GetByID(ctx context.Context, id string) (*Match, error) {
	m, ok := r.Matches[id]
	if !ok {
		return nil, ErrMatchNotFound
	}
	return m, nil
}
