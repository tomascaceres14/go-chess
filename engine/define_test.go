package gochess

import (
	"testing"
)

func newTestStartingPos() *Game {
	game, _ := NewGameClassic("Player_White", "Player_Black")
	return game
}

func newTestFENPos(pos string) *Game {
	game, _ := NewGameFENString(pos, "Player_White", "Player_Black")
	return game
}

func TestNewClassicGame(t *testing.T) {
	_, err := NewGameClassic("Player_White", "Player_Black")
	if err != nil {
		t.Errorf(`NewGame() has error %v`, err)
	}
}

func TestFENStringInitialPos(t *testing.T) {
	FENString := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	g, err := NewGameFENString(FENString, "Player_White", "Player_Black")
	if err != nil {
		t.Errorf("Error creating game from FENString: %s", err)
	}

	gameFENString := g.GetFENString()
	if gameFENString != FENString {
		t.Errorf(`NewGameFENString(). Got %s want %s`, gameFENString, FENString)
	}
}

func TestPawnStartingForwardCantJump(t *testing.T) {
	FENString := "rnbqkbnr/pppp1ppp/8/4p3/8/P7/1PPPPPPP/RNBQKBNR b KQkq - 0 1"
	testName := "TestPawnStartingForwardCantJump"

	game, err := NewGameFENString(FENString, "Player_White", "Player_Black")
	if err != nil {
		t.Errorf("Error creating game from FENString: %s", err)
	}

	from := "e5"
	to := "e3"
	movesWhite := false

	if err := game.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}
}
