package gochess

import "testing"

func TestCantGoOutofBounds(t *testing.T) {
	testName := "TestCantGoOutofBounds"
	engine := testStartingPos()

	from := "a1"
	to := "i1"
	movesWhite := true

	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	pos := "rnbqkbn1/ppppppp1/8/8/8/8/PPPPPPP1/RNBQKBNR w KQq - 0 1"
	engine = testFENPos(pos)

	from = "a1"
	to = "h9"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}
}

func TestPinnedPieceCantMove(t *testing.T) {
	testName := "TestPinnedPieceCantMove"
	pos := "rnbqk1nr/pppp1ppp/8/4p3/1b1P4/2N5/PPP1PPPP/R1BQKBNR w KQkq - 2 3"
	engine := testFENPos(pos)

	from := "c3"
	to := "d5"
	movesWhite := true

	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}
}
