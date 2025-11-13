package gochess

import (
	"fmt"
	"testing"
)

func TestKingCantCaptureOwnPieces(t *testing.T) {
	testName := "TestKingCantCaptureOwnPieces"
	engine := newTestStartingPos()

	from := "e1"
	to := "e2"
	movesWhite := true

	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

}

func TestKingCanMove(t *testing.T) {
	testName := "TestKingCantCaptureOwnPieces"
	engine := newTestStartingPos()

	from := "e2"
	to := "e4"
	movesWhite := true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	from = "d7"
	to = "d5"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	from = "e1"
	to = "e2"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	from = "e8"
	to = "d7"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}
	fmt.Println(engine.game.gameBoard)
}

func TestKingCanCastles(t *testing.T) {
	testName := "TestKingCanCastle"
	engine := newTestFENPos("r2qkbnr/ppp1pppp/2n5/3p1b2/2B1P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 4 4")

	from := "e1"
	to := "g1"
	movesWhite := true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	from = "d8"
	to = "d7"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	from = "a2"
	to = "a3"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	from = "e8"
	to = "c8"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	fmt.Println(engine.game.gameBoard)
}
