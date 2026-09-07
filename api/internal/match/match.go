package match

import (
	"context"
	"log"
	"sync"

	gochess "github.com/tomascaceres14/go-chess/engine"
)

var (
	StatusPending     = "PENDING"
	StatusMatchmaking = "MATCHMAKING"
	StatusPlaying     = "PLAYING"
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
	OwnerWhite  bool
	listeners   map[string]chan GameResponse
	CommandsCh  chan GameCommand
	MoveHistory []string
	mu          sync.RWMutex
	FEN         string
}

type Repository interface {
	Save(ctx context.Context, match *Match) (*Match, error)
	GetByID(ctx context.Context, id string) (*Match, error)
	GetMatchesByUserID(ctx context.Context, id string) ([]*Match, error)
	FinalizeMatch(ctx context.Context, matchID, status, FEN string, moveHistory []string) error
	SetStatus(ctx context.Context, matchID, status string) error
	SetStatusAndOpponent(ctx context.Context, matchID, opponentID, status string) error
}

func NewMatch(userID string, color bool) *Match {
	return &Match{
		OwnerID:     userID,
		OpponentID:  "",
		OwnerWhite:  color,
		Status:      StatusPending,
		listeners:   make(map[string]chan GameResponse),
		CommandsCh:  make(chan GameCommand, 2),
		MoveHistory: make([]string, 0),
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

	white := m.OwnerID
	black := m.OpponentID

	if !m.OwnerWhite {
		white = m.OpponentID
		black = m.OwnerID
	}

	// Ignoring error until refactoring. No need to provide white and black ids or names
	//game, _ := gochess.NewGameClassic(white, black)
	game, _ := gochess.NewGameFENString("rnbqkbnr/pppp1ppp/8/4p3/5PP1/8/PPPPP2P/RNBQKBNR b KQkq g3 0 2", white, black)
	m.FEN = game.GetFENString()

	defer m.CloseChannels()

	// Send game begin signal
	m.sendMessage(GameResponse{
		Command: MatchBeginStatus,
		Valid:   true,
		Fen:     game.GetFENString(),
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
					Status:  m.Status,
					Valid:   false,
				}

				// Execute move
				err := game.MovePlayer(msg.Move.From, msg.Move.To, msg.UserID)

				// Send error if any
				if err != nil {
					log.Println(err)
					response.Error = err.Error()
					m.sendMessage(response)
					continue
				}

				// Refresh game grid
				response.UserID = ""
				response.Valid = true
				response.Fen = game.GetFENString()

				if game.Status() != gochess.StatusPlaying {
					response.Command = MatchEndCmd
					response.Status = game.Status()
					m.MoveHistory = game.MoveHistory()
					m.sendMessage(response)
					return
				}

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
