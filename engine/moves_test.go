package gochess

import (
	"testing"
)

func TestCantGoOutofBounds(t *testing.T) {
	testName := "TestCantGoOutofBounds"
	engine := newTestStartingPos()

	from := "a1"
	to := "i1"
	movesWhite := true

	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	pos := "rnbqkbn1/ppppppp1/8/8/8/8/PPPPPPP1/RNBQKBNR w KQq - 0 1"
	engine = newTestFENPos(pos)

	from = "a1"
	to = "h9"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}
}

func TestPinnedPieceCantMove(t *testing.T) {
	testName := "TestPinnedPieceCantMove"
	pos := "rnbqk1nr/pppp1ppp/8/4p3/1b1P4/2N5/PPP1PPPP/R1BQKBNR w KQkq - 2 3"
	engine := newTestFENPos(pos)

	from := "c3"
	to := "d5"
	movesWhite := true

	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}
}

func TestPinnedPieceCanCaptureAttacker(t *testing.T) {
	testName := "TestPinnedPieceCantMove"
	pos := "rnbqkbnr/ppp1pppp/8/8/8/8/PPPQPPPP/RNBK1BNR w kq - 0 1"
	engine := newTestFENPos(pos)

	// Pinned piece cant move out of pin
	from := "d2"
	to := "c3"
	movesWhite := true
	if _, err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	// Pinned piece can move in same line as attacker
	from = "d2"
	to = "d4"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	from = "a7"
	to = "a6"
	movesWhite = false
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	// Pinned piece can take attacker
	from = "d4"
	to = "d8"
	movesWhite = true
	if _, err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}
}
