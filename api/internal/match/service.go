package match

import (
	"context"
	"errors"
)

var (
	ErrMatchAlreadyExists   = errors.New("Match already exists")
	ErrMatchNotFound        = errors.New("Match not found")
	ErrOwnerNotConnected    = errors.New("Owner not yet connected")
	ErrMatchFull            = errors.New("Match is full")
	ErrUserAlreadyConnected = errors.New("User already connected")
)

type Service struct {
	repo         Repository
	matchManager *MatchManager
}

func NewService(r Repository) *Service {
	return &Service{
		repo:         r,
		matchManager: NewMatchmaker(),
	}
}

func (s *Service) StartNewMatch(ctx context.Context, userID string, whites bool) (*Match, error) {
	match := NewMatch(userID, whites)
	if err := s.matchManager.AddMatch(match); err != nil {
		return nil, err
	}
	return match, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*Match, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) AddPlayerToMatch(ctx context.Context, matchID, userID string) (chan GameResponse, error) {
	match, err := s.matchManager.GetMatch(matchID)
	if err != nil {
		return nil, err
	}

	switch match.Status {
	case StatusPending:
		if userID == match.OwnerID {
			s.matchManager.SetStatus(matchID, StatusMatchmaking)
		} else {
			return nil, ErrOwnerNotConnected
		}
	case StatusMatchmaking:
		s.matchManager.SetOpponentID(matchID, userID)
		if err := match.Start(); err != nil {
			return nil, err
		}
	}

	ch, err := s.matchManager.AddListener(matchID, userID)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (s *Service) GetCommandsCh(gameID, userID string) (chan GameCommand, error) {
	match, err := s.matchManager.GetMatch(gameID)
	if err != nil {
		return nil, err
	}

	return match.CommandsCh, nil
}
