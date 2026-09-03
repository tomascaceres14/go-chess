package match

import (
	"context"
	"errors"
	"log"
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

	match, err := s.repo.Save(ctx, NewMatch(userID, whites))
	if err != nil {
		return nil, err
	}

	if err := s.matchManager.AddMatch(match); err != nil {
		return nil, err
	}
	return match, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*Match, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) AddUserToMatch(ctx context.Context, matchID, userID string) (chan GameResponse, error) {
	match, err := s.matchManager.GetMatch(matchID)
	if err != nil {
		return nil, err
	}

	switch match.Status {
	case StatusPending:
		if userID == match.OwnerID {
			log.Printf("Owner for match %s connected. Changing status", matchID)
			s.matchManager.SetStatus(matchID, StatusMatchmaking)
			s.repo.SetStatus(ctx, matchID, StatusMatchmaking)
		} else {
			return nil, ErrOwnerNotConnected
		}

	case StatusMatchmaking:
		log.Printf("Opponent found for match %s.", matchID)
		s.matchManager.SetOpponentID(matchID, userID)
		if err := s.repo.SetStatusAndOpponent(ctx, matchID, userID, StatusPlaying); err != nil {
			return nil, err
		}

		go match.Start()
	}

	ch, err := s.matchManager.AddListener(matchID, userID)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (s *Service) RemoveUserFromMatch(matchID, userID string) {
	s.matchManager.RemoveListener(matchID, userID)
}

func (s *Service) GetCommandsCh(gameID, userID string) (chan GameCommand, error) {
	match, err := s.matchManager.GetMatch(gameID)
	if err != nil {
		return nil, err
	}

	return match.CommandsCh, nil
}

func (s *Service) GetMatchesByUserID(ctx context.Context, userID string) ([]*Match, error) {
	return s.repo.GetMatchesByUserID(ctx, userID)
}

func (s *Service) FinalizeMatch(ctx context.Context, matchID, status, FEN string, moveHistory []string) error {
	return s.repo.FinalizeMatch(ctx, matchID, status, FEN, moveHistory)
}
