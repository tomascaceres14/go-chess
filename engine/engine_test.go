package gochess

import (
	"fmt"
	"testing"
)

func newTestStartingPos() ChessEngine {
	engine := NewChessEngine()

	whiteName := "Player_White"
	blackName := "Player_Black"

	engine.NewGame(whiteName, blackName)

	return *engine
}

func newTestFENPos(pos string) ChessEngine {
	engine := NewChessEngine()

	whiteName := "Player_White"
	blackName := "Player_Black"

	engine.NewGameFENString(pos, whiteName, blackName)

	return *engine
}

func TestNewClassicGame(t *testing.T) {
	engine := NewChessEngine()

	whiteName := "Player_White"
	blackName := "Player_Black"

	_, err := engine.NewGame(whiteName, blackName)
	if err != nil {
		t.Errorf(`NewGame() has error %v`, err)
	}
}

func TestFENStringInitialPos(t *testing.T) {
	FENString := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	engine := NewChessEngine()

	_, err := engine.NewGameFENString(FENString, "Player_White", "Player_Black")
	if err != nil {
		t.Errorf("Error creating game from FENString: %s", err)
	}

	engineFENString := engine.GetFENString()

	if engineFENString != FENString {
		t.Errorf(`NewGameFENString(). Got %s want %s`, engineFENString, FENString)
	}
}

func TestPawnStartingForwardCantJump(t *testing.T) {
	FENString := "rnbqkbnr/pppp1ppp/8/4p3/8/P7/1PPPPPPP/RNBQKBNR b KQkq - 0 1"
	testName := "TestPawnStartingForwardCantJump"

	engine := NewChessEngine()

	_, err := engine.NewGameFENString(FENString, "Player_White", "Player_Black")
	if err != nil {
		t.Errorf("Error creating game from FENString: %s", err)
	}

	from := "e5"
	to := "e3"
	movesWhite := false
	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	fmt.Println(engine.game.gameBoard)
}
