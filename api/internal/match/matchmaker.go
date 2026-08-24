package match

import (
	"sync"
)

type MatchMaker struct {
	matches map[string]*Match
	mu      sync.RWMutex
}

func NewMatchmaker() *MatchMaker {
	return &MatchMaker{
		matches: make(map[string]*Match),
	}
}

func (mm *MatchMaker) Add(match *Match) error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	_, ok := mm.matches[match.ID]
	if ok {
		return ErrMatchAlreadyExists
	}

	mm.matches[match.ID] = match

	return nil
}

func (mm *MatchMaker) Get(id string) (*Match, error) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	match, ok := mm.matches[id]
	if !ok {
		return nil, ErrMatchNotFound
	}
	return match, nil
}

func (mm *MatchMaker) UpdateStatus(id string, status string) (*Match, bool) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	match, ok := mm.matches[id]
	match.Status = status

	return match, ok
}

func (mm *MatchMaker) AssignOpponent(id, userID string) error {
	match, err := mm.Get(id)
	if err != nil {
		return err
	}

	match.OpponentID = userID

	mm.matches[id] = match
	return nil
}
