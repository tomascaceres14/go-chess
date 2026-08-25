package match

import gochess "github.com/tomascaceres14/go-chess/engine"

const (
	MovePieceCmd = "match.cmd.move"
)

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
	UserID string        `json:"user_id"`
	Cmd    string        `json:"cmd"`
	Valid  bool          `json:"valid"`
	Error  error         `json:"error,omitempty"`
	Grid   *gochess.Grid `json:"grid"`
}
