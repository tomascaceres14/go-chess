package match

import (
	"context"
	"errors"
)

var (
	ErrMatchAlreadyExists = errors.New("Match already exists")
	ErrMatchNotFound      = errors.New("Match not found")
)

type Service struct {
	repo       Repository
	matchMaker *MatchMaker
}

func NewService(r Repository) *Service {
	return &Service{
		repo:       r,
		matchMaker: NewMatchmaker(),
	}
}

func (s *Service) StartNewMatch(ctx context.Context, userID string, whites bool) (*Match, error) {
	match := NewMatch(userID, whites)
	if err := s.matchMaker.Add(match); err != nil {
		return nil, err
	}
	return match, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*Match, error) {
	return s.repo.GetByID(ctx, id)
}
