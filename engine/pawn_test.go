package gochess

import (
	"fmt"
	"testing"
)

func TestPawnMoveForward(t *testing.T) {
	testName := "TestPawnMoveForward"
	engine := newTestStartingPos()

	from := "e2"
	to := "e3"
	movesWhite := true

	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

}

func TestPawnJump(t *testing.T) {
	testName := "TestPawnJump"
	engine := newTestStartingPos()

	from := "e2"
	to := "e4"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "a7"
	to = "a5"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e4"
	to = "e6"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}
}

func TestPawnCantCaptureForward(t *testing.T) {
	testName := "TestPawnCantCaptureForward"
	engine := newTestStartingPos()

	from := "e2"
	to := "e4"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e7"
	to = "e5"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e4"
	to = "e5"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

}

func TestPawnEnPassantLeftOption(t *testing.T) {
	testName := "TestPawnEnPassantLeftOption"
	engine := newTestStartingPos()

	from := "e2"
	to := "e4"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "a7"
	to = "a6"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e4"
	to = "e5"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "d7"
	to = "d5"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	fmt.Println(engine.game.gameBoard)

	from = "e5"
	to = "d6"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

}

func TestPawnEnPassantRightOption(t *testing.T) {
	testName := "TestPawnEnPassantRightOption"
	engine := newTestStartingPos()

	from := "e2"
	to := "e4"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "a7"
	to = "a5"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e4"
	to = "e5"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "f7"
	to = "f5"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e5"
	to = "f6"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

}

func TestPawnEnPassantSandwitched(t *testing.T) {
	testName := "TestPawnEnPassantRightOption"
	engine := newTestFENPos("rnbqkbnr/1pp1p1pp/p7/3pPp2/3P4/8/PPP2PPP/RNBQKBNR w KQkq f6 0 4")

	from := "e5"
	to := "d6"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e5"
	to = "f6"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

}

func TestPawnEnPassantFEN(t *testing.T) {
	testName := "TestPawnEnPassantFEN"
	pos := "rnbqkbnr/1pp1pppp/p7/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3"
	engine := newTestFENPos(pos)

	from := "e5"
	to := "d6"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}
}

func TestPawnCantEnPassantNextTurn(t *testing.T) {
	testName := "TestPawnCantEnPassantNextTurn"
	pos := "rnbqkbnr/1pp1pppp/p7/3pP3/8/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 3"
	engine := newTestFENPos(pos)

	from := "a2"
	to := "a3"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "h7"
	to = "h6"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "e5"
	to = "d6"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

}

func TestPawnCantEnPassantOtherPieces(t *testing.T) {
	testName := "TestPawnCantEnPassantOtherPieces"
	pos := "r1bqkbnr/pppppppp/8/3nP3/8/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"
	engine := newTestFENPos(pos)

	from := "e5"
	to := "d6"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

}
