package match

import gochess "github.com/tomascaceres14/go-chess/engine"

const (
	MovePieceCmd = "match.cmd.move"

	MatchBeginStatus   = "match.status.begin"
	MatchWaitingStatus = "match.status.waiting"
)

type NewMatchParams struct {
	UserID string `json:"user_id"`
	Whites bool   `json:"whites"`
}

type Move struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type GameCommand struct {
	UserID  string `json:"-"`
	Command string `json:"cmd"`
	Move    Move   `json:"move"`
}

type GameResponse struct {
	UserID  string        `json:"user_id"`
	Command string        `json:"cmd"`
	Valid   bool          `json:"valid"`
	Error   error         `json:"error,omitempty"`
	Grid    *gochess.Grid `json:"grid"`
}
