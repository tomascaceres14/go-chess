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

func (s *Service) SetStatus(ctx context.Context, matchID, status string) error {
	if err := s.matchManager.SetStatus(matchID, status); err != nil {
		return err
	}

	if err := s.repo.SetStatus(ctx, matchID, status); err != nil {
		return err
	}

	return nil
}

func (s *Service) AssignOwner(ctx context.Context, matchID, userID string) error {
	m, err := s.matchManager.GetMatch(matchID)
	if err != nil {
		return err
	}

	m.mu.RLock()
	if userID != m.WhitesID {
		return ErrOwnerNotConnected
	}
	m.mu.RUnlock()

	log.Printf("Owner for match %s connected. Changing status", matchID)
	if err := s.SetStatus(ctx, matchID, StatusMatchmaking); err != nil {
		return err
	}

	return nil
}

func (s *Service) AssignOpponent(ctx context.Context, matchID, userID string) error {
	if err := s.repo.SetStatusAndOpponent(ctx, matchID, userID, StatusPlaying); err != nil {
		return err
	}

	if err := s.matchManager.SetOpponentID(matchID, userID); err != nil {
		return err
	}

	if err := s.matchManager.SetStatus(matchID, StatusPlaying); err != nil {
		return err
	}

	return nil
}

func (s *Service) AddUserToMatch(ctx context.Context, matchID, userID string) (chan GameResponse, error) {
	log.Printf("Player %s attempting connection to match %s", userID, matchID)
	status, err := s.matchManager.GetStatus(matchID)
	if err != nil {
		return nil, err
	}

	switch status {
	case StatusPending:
		if err := s.AssignOwner(ctx, matchID, userID); err != nil {
			return nil, err
		}
	case StatusMatchmaking:
		log.Printf("Opponent found for match %s.", matchID)
		if err := s.AssignOpponent(ctx, matchID, userID); err != nil {
			return nil, err
		}
		go s.StartMatch(matchID)
	}

	ch, err := s.matchManager.AddListener(matchID, userID)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (s *Service) StartMatch(matchID string) error {
	match, err := s.matchManager.GetMatch(matchID)
	if err != nil {
		return err
	}

	go match.Start()

	return nil
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

func (s *Service) FinalizeMatch(ctx context.Context, matchID, status, FEN string) error {

	// Retrieve match
	m, err := s.matchManager.GetMatch(matchID)
	if err != nil {
		return err
	}

	return s.repo.FinalizeMatch(ctx, matchID, status, FEN, m.MoveHistory)
}

func (s *Service) GetLiveMatches(ctx context.Context) []*Match {
	return s.matchManager.GetMatches()
}
