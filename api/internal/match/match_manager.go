package match

import (
	"log"
	"maps"
	"slices"
	"sync"
)

type MatchManager struct {
	matches map[string]*Match
	mu      sync.RWMutex
}

func NewMatchmaker() *MatchManager {
	return &MatchManager{
		matches: make(map[string]*Match),
	}
}

func (mm *MatchManager) AddMatch(match *Match) error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	_, ok := mm.matches[match.ID]
	if ok {
		return ErrMatchAlreadyExists
	}

	mm.matches[match.ID] = match

	return nil
}

func (mm *MatchManager) GetMatch(id string) (*Match, error) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	match, ok := mm.matches[id]
	if !ok {
		return nil, ErrMatchNotFound
	}
	return match, nil
}

func (mm *MatchManager) GetStatus(id string) (string, error) {
	m, err := mm.GetMatch(id)
	if err != nil {
		return "", err
	}

	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.Status, nil
}

func (mm *MatchManager) SetStatus(id string, status string) error {
	mm.mu.RLock()
	match, err := mm.GetMatch(id)
	if err != nil {
		return err
	}
	mm.mu.RUnlock()

	match.mu.Lock()
	match.Status = status
	match.mu.Unlock()
	return nil
}

func (mm *MatchManager) SetOpponentID(id, userID string) error {
	match, err := mm.GetMatch(id)
	if err != nil {
		return err
	}

	match.BlacksID = userID

	mm.matches[id] = match
	return nil
}

func (mm *MatchManager) GetListener(matchID, userID string) (chan GameResponse, error) {
	log.Printf("Getting listener %s for match %s", userID, matchID)
	match, err := mm.GetMatch(matchID)
	if err != nil {
		return nil, err
	}
	ch, ok := match.listeners[userID]
	if !ok {
		return nil, ErrUserAlreadyConnected
	}

	return ch, nil
}

func (mm *MatchManager) AddListener(matchID, userID string) (chan GameResponse, error) {
	log.Printf("Adding listener %s for match %s", userID, matchID)
	match, err := mm.GetMatch(matchID)
	if err != nil {
		return nil, err
	}
	_, ok := match.listeners[userID]
	if ok {
		return nil, ErrUserAlreadyConnected
	}

	ch := make(chan GameResponse, 1)
	match.listeners[userID] = ch
	return ch, nil
}

func (mm *MatchManager) RemoveListener(matchID, userID string) {
	log.Printf("Removing listener %s for match %s", userID, matchID)

	match, err := mm.GetMatch(matchID)
	if err != nil {
		return
	}

	match.RemoveListener(userID)
}

func (mm *MatchManager) GetMatches() []*Match {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	return slices.Collect(maps.Values(mm.matches))
}
