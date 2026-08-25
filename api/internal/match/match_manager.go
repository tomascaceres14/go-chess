package match

import (
	"log"
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

func (mm *MatchManager) SetStatus(id string, status string) (*Match, bool) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	match, ok := mm.matches[id]
	match.Status = status

	return match, ok
}

func (mm *MatchManager) SetOpponentID(id, userID string) error {
	match, err := mm.GetMatch(id)
	if err != nil {
		return err
	}

	match.OpponentID = userID

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
	log.Println("GetListener listener ch", ch, ok, userID)
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
	listener, ok := match.listeners[userID]
	log.Println("AddListener listener ch", listener, ok, userID)
	if ok {
		return nil, ErrUserAlreadyConnected
	}

	ch := make(chan GameResponse, 1)
	match.listeners[userID] = ch
	return ch, nil
}
