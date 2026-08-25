package match

import (
	"context"
	"errors"
)

var (
	ErrMatchAlreadyExists = errors.New("Match already exists")
	ErrMatchNotFound      = errors.New("Match not found")
	ErrOwnerNotConnected  = errors.New("Owner not yet connected")
	ErrMatchFull          = errors.New("Match is full")
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
	if err := s.matchManager.Add(match); err != nil {
		return nil, err
	}
	return match, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*Match, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) AssignPlayerToMatch(ctx context.Context, gameID, userID string) (*Match, error) {
	match, err := s.matchManager.Get(gameID)
	if err != nil {
		return nil, err
	}

	switch match.Status {
	case StatusPending:
		if userID == match.OwnerID {
			s.matchManager.SetStatus(gameID, StatusMatchmaking)
		} else {
			return nil, ErrOwnerNotConnected
		}
	case StatusMatchmaking:
		s.matchManager.SetOpponentID(gameID, userID)
		if err := match.Start(ctx); err != nil {
			return nil, err
		}
	default:
		return nil, ErrMatchFull
	}

	return match, nil
}
