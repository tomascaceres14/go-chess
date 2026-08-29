package match

import (
	"context"
	"log"
	"sync"

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
	OwnerID     string
	OpponentID  string
	Status      string
	Result      string
	OwnerWhite  bool
	listeners   map[string]chan GameResponse
	CommandsCh  chan GameCommand
	MoveHistory []gochess.Move
	mu          sync.RWMutex
}

type Repository interface {
	Save(ctx context.Context, match *Match) (*Match, error)
	GetByID(ctx context.Context, id string) (*Match, error)
}

func NewMatch(userID string, color bool) *Match {
	return &Match{
		ID:          uuid.NewString(),
		OwnerID:     userID,
		OpponentID:  "",
		OwnerWhite:  color,
		Status:      StatusPending,
		listeners:   make(map[string]chan GameResponse),
		CommandsCh:  make(chan GameCommand, 2),
		MoveHistory: make([]gochess.Move, 0),
	}
}

func (m *Match) RemoveListener(userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	ch, ok := m.listeners[userID]
	if !ok {
		return
	}

	close(ch)
	delete(m.listeners, userID)
}

func (m *Match) Start() {

	// Define players colors. For now I'll just leave it like this, although
	// I'm not sure if the engine should receive player's names and assign them to each team or
	// just initialize the game and let the server handle the user->color relation
	white := m.OwnerID
	black := m.OpponentID

	if !m.OwnerWhite {
		white = m.OpponentID
		black = m.OwnerID
	}

	// Ignoring error until refactoring. No need to provide white and black ids or names
	game, _ := gochess.NewGameClassic(white, black)

	defer m.CloseChannels()

	// Send game begin signal
	m.sendMessage(GameResponse{
		Command: MatchBeginStatus,
		Valid:   true,
		Grid:    game.GetFlattenString(),
	})

	for {
		select {
		// Recieve user messages
		case msg := <-m.CommandsCh:
			switch msg.Command {
			case MovePieceCmd:

				// Prepare initial message
				response := GameResponse{
					UserID:  msg.UserID,
					Command: msg.Command,
					Valid:   false,
				}

				// Execute move
				// TODO: Adapt Move() func to receive player name and validate based on that,
				// not on color trying to move.
				err := game.MovePlayer(msg.Move.From, msg.Move.To, msg.UserID)

				// Adjust response based on error
				if err != nil {
					log.Println(err)
					response.Error = err.Error()
				} else {
					response.UserID = ""
					response.Valid = true
				}

				// Refresh game grid and respond
				response.Grid = game.GetFlattenString()
				m.sendMessage(response)
			}
		case <-context.Background().Done():

		}
	}
}

func (m *Match) fanout(msg GameResponse) {
	for _, ch := range m.listeners {
		ch <- msg
	}
}

func (m *Match) sendMessage(msg GameResponse) {
	if msg.UserID == "" {
		m.fanout(msg)
		return
	}
	m.listeners[msg.UserID] <- msg
}

func (m *Match) CloseChannels() {
	for _, ch := range m.listeners {
		close(ch)
	}
}
