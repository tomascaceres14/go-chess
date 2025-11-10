package gochess

import (
	"fmt"
	"testing"
)

func TestPawnMoveForward(t *testing.T) {
	testName := "TestPawnMoveForward"
	engine := testStartingPos()

	from := "e2"
	to := "e3"
	movesWhite := true

	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

}

func TestPawnJump(t *testing.T) {
	testName := "TestPawnJump"
	engine := testStartingPos()

	from := "e2"
	to := "e4"
	movesWhite := true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "a7"
	to = "a5"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e4"
	to = "e6"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}
}

func TestPawnCantCaptureForward(t *testing.T) {
	testName := "TestPawnCantCaptureForward"
	engine := testStartingPos()

	from := "e2"
	to := "e4"
	movesWhite := true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e7"
	to = "e5"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e4"
	to = "e5"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

}

func TestPawnEnPassant(t *testing.T) {
	testName := "TestPawnEnPassant"
	engine := testStartingPos()

	from := "e2"
	to := "e4"
	movesWhite := true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "a7"
	to = "a6"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e4"
	to = "e5"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "d7"
	to = "d5"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e5"
	to = "d6"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	fmt.Println(engine.game.gameBoard)
}

func TestPawnEnPassantFEN(t *testing.T) {
	testName := "TestPawnEnPassantFEN"
	pos := "rnbqkbnr/1pp1pppp/p7/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3"
	engine := testFENPos(pos)

	from := "e5"
	to := "d6"
	movesWhite := true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}
}

func TestPawnCantEnPassantAfterMove(t *testing.T) {
	testName := "TestPawnCantEnPassantAfterMove"
	pos := "rnbqkbnr/1pp1pppp/p7/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3"
	engine := testFENPos(pos)

	from := "a2"
	to := "a3"
	movesWhite := true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "h7"
	to = "h6"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e5"
	to = "d6"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	fmt.Println(engine.game.gameBoard)

}
