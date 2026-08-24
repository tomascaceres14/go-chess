package match

import (
	"context"

	"github.com/google/uuid"
	gochess "github.com/tomascaceres14/go-chess/engine"
)

var (
	StatusPending     = "PENDING"
	StatusMatchmaking = "MATCHMAKING"
	StatusOngoing     = "ONGOING"
	StatusAborted     = "ABORTED"
	StatusDraw        = "DRAW"
	StatusWhiteWins   = "WHITE_WINS"
	StatusBlackWins   = "BLACK_WINS"
)

type Match struct {
	ID          string
	OwnerColour bool
	OwnerID     string
	OpponentID  string
	Status      string
	Result      string
	MoveHistory []gochess.Move
}

type Repository interface {
	Save(ctx context.Context, match *Match) (*Match, error)
	FindByID(ctx context.Context, id string) (*Match, error)
	GetByID(ctx context.Context, id string) (*Match, error)
}

func NewMatch(userID string, colour bool) *Match {

	whiteID := userID
	blackID := ""

	if !colour {
		whiteID = ""
		blackID = userID
	}

	return &Match{
		ID:          uuid.NewString(),
		OwnerID:     whiteID,
		OpponentID:  blackID,
		OwnerColour: colour,
		Status:      StatusPending,
		MoveHistory: make([]gochess.Move, 0),
	}
}
