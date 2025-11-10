package gochess

import "testing"

func testStartingPos() ChessEngine {
	engine := NewChessEngine()

	whiteName := "Player_White"
	blackName := "Player_Black"

	engine.NewGame(whiteName, blackName)

	return *engine
}

func testFENPos(pos string) ChessEngine {
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
