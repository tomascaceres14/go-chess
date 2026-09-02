package match

import (
	"context"

	"github.com/google/uuid"
	"github.com/tomascaceres14/go-chess/api/internal/database/generated"
)

// For the sake of simplicity, PostgresRepository will just use the default
// *sqlc.Queries, even though it probably should have it's own interface to avoid
// implementing other model's methods.
type PostgresRepository struct {
	db *generated.Queries
}

func NewPostgresRepository(db *generated.Queries) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Save(ctx context.Context, m *Match) (*Match, error) {
	saveMoveHistory := make([]string, len(m.MoveHistory))
	for i, v := range m.MoveHistory {
		saveMoveHistory[i] = v.String()
	}
	ownerID, err := uuid.Parse(m.OwnerID)
	if err != nil {
		return nil, err
	}

	opponentID, err := uuid.Parse(m.OpponentID)
	if err != nil {
		return nil, err
	}

	id, err := r.db.CreateMatch(ctx, generated.CreateMatchParams{

		OwnerID:     ownerID,
		OpponentID:  opponentID,
		Status:      m.Status,
		Result:      m.Result,
		OwnerWhite:  m.OwnerWhite,
		MoveHistory: saveMoveHistory,
	})

	if err != nil {
		return nil, err
	}

	m.ID = id.String()

	return m, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*Match, error) {
	matchID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	m, err := r.db.GetMatchByID(ctx, matchID)
	if err != nil {
		return nil, err
	}

	return &Match{
		ID:          m.ID.String(),
		OwnerID:     m.OwnerID.String(),
		OpponentID:  m.OpponentID.String(),
		Status:      m.Status,
		Result:      m.Result,
		OwnerWhite:  m.OwnerWhite,
		MoveHistory: nil,
	}, nil
}
