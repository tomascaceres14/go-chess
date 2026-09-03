package match

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

	ownerID, err := uuid.Parse(m.OwnerID)
	if err != nil {
		return nil, err
	}

	id, err := r.db.CreateMatch(ctx, generated.CreateMatchParams{
		OwnerID:     ownerID,
		Status:      m.Status,
		OwnerWhite:  m.OwnerWhite,
		MoveHistory: m.MoveHistory,
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
		OwnerWhite:  m.OwnerWhite,
		MoveHistory: nil,
	}, nil
}

func (r *PostgresRepository) GetMatchesByUserID(ctx context.Context, id string) ([]*Match, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	matchesDB, err := r.db.GetMatchesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return ParseMatchDBList(matchesDB), nil
}

func (r *PostgresRepository) FinalizeMatch(ctx context.Context, matchID, status, FEN string, moveHistory []string) error {
	id, err := uuid.Parse(matchID)
	if err != nil {
		return err
	}

	return r.db.UpdateGameFinalState(ctx, generated.UpdateGameFinalStateParams{
		ID:     id,
		Status: status,
		Fen: pgtype.Text{
			String: FEN,
			Valid:  true,
		},
		MoveHistory: moveHistory,
	})
}

func (r *PostgresRepository) SetStatus(ctx context.Context, matchID, status string) error {
	id, err := uuid.Parse(matchID)
	if err != nil {
		return err
	}
	return r.db.SetMatchStatus(ctx, generated.SetMatchStatusParams{
		ID:     id,
		Status: status,
	})
}

func (r *PostgresRepository) SetStatusAndOpponent(ctx context.Context, matchID, opponentID, status string) error {
	id, err := uuid.Parse(matchID)
	if err != nil {
		return err
	}
	opponent, err := uuid.Parse(opponentID)
	if err != nil {
		return err
	}
	return r.db.SetMatchStatusAndOpponent(ctx, generated.SetMatchStatusAndOpponentParams{
		ID: id,
		OpponentID: pgtype.UUID{ // ? Should be uuid.UUID
			Bytes: opponent,
			Valid: true,
		},
		Status: status,
	})
}

func ParseMatchDB(m *generated.Match) *Match {
	return &Match{
		ID:         m.ID.String(),
		OwnerID:    m.OwnerID.String(),
		OpponentID: m.OpponentID.String(),
		Status:     m.Status,
		OwnerWhite: m.OwnerWhite,
	}
}

func ParseMatchDBList(matches []generated.Match) []*Match {
	parsedMatches := make([]*Match, len(matches))

	for i, v := range matches {
		parsedMatches[i] = ParseMatchDB(&v)
	}

	return parsedMatches
}
