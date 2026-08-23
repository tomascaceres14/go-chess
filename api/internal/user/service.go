package user

import "context"

type Service struct {
	repo Repository
}

func (s *Service) Register(ctx context.Context, register RegisterUser) (*User, error) {
	if err := register.Validate(); err != nil {
		return nil, err
	}

	if register.Password != register.RepeatPassword {
		return nil, ErrPasswordsDontMatch
	}

	user, err := NewUser(register.Username, register.Password)
	if err != nil {
		return nil, err
	}

	s.repo.Save(ctx, user)
	return nil, nil
}
