package match

import gochess "github.com/tomascaceres14/go-chess/engine"

type GameCommand struct {
	UserID   string
	Cmd      string
	MovePlay string
	Game     *gochess.Game
}

type GameResponse struct {
	UserID string
	Cmd    string
	Valid  bool
	Error  error
	Grid   *gochess.Grid
}
