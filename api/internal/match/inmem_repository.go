package match

import (
	"context"
	"maps"
	"slices"
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

func (r *MemoryRepository) GetAll(ctx context.Context) []*Match {
	matches := slices.Collect(maps.Values(r.Matches))
	if len(matches) == 0 {
		return []*Match{}
	}

	return matches
}
