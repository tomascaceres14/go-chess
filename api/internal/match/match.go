package match

import (
	"context"
	"log"

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
	OwnerWhite  bool
	OwnerID     string
	OpponentID  string
	Status      string
	Result      string
	CommandsCh  chan GameCommand
	ResponsesCh chan GameResponse
	MoveHistory []gochess.Move
}

type Repository interface {
	Save(ctx context.Context, match *Match) (*Match, error)
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
		OwnerWhite:  colour,
		Status:      StatusPending,
		MoveHistory: make([]gochess.Move, 0),
	}
}

func (m *Match) Start(ctx context.Context) error {
	white := m.OwnerID
	black := m.OpponentID

	if !m.OwnerWhite {
		white = m.OpponentID
		black = m.OwnerID
	}
	game, err := gochess.NewGameClassic(white, black)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Graceful shutdown of match %s", m.ID)
				break
			case msg := <-m.CommandsCh:
				switch msg.Command {
				case MovePieceCmd:
					color := m.OwnerWhite
					if msg.UserID == m.OpponentID {
						color = !color
					}

					err := game.Move(msg.Move.From, msg.Move.To, color)

					// Nil error, movement is approved. Empty userID means send to both players
					if err != nil {
						log.Println(err)
					} else {
						msg.UserID = ""
					}

					response := &GameResponse{
						UserID: msg.UserID,
						Cmd:    msg.Command,
						Valid:  err != nil,
						Error:  err,
					}
				}
			}
		}
	}()

	return nil
}
