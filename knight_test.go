package gochess

import "testing"

func TestKnightForks(t *testing.T) {
	testName := "TestPinnedPieceCantMove"
	pos := "r1q1k1nr/pp1pp2p/8/8/2N1N3/8/PPPPPPPP/R1BQKB1R w KQkq - 0 1"
	engine := newTestFENPos(pos)

	// Enter fork
	from := "c4"
	to := "d6"
	movesWhite := true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	// Cant move piece that isn't king or capturing knight
	from = "c8"
	to = "b8"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	// Take forking knight
	from = "e7"
	to = "d6"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	// Enter fork
	from = "e4"
	to = "d6"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	// Cant move piece that isn't king. Same test but now cant capture knight
	from = "c8"
	to = "b8"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	// Move king
	from = "e8"
	to = "d8"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}

	// Capture queen
	from = "e8"
	to = "d8"
	movesWhite = false
	if err := engine.Move(from, to, movesWhite); err == nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err, got %v", testName, from, to, movesWhite, err)
	}

	// Enter fork
	from = "d6"
	to = "c8"
	movesWhite = true
	if err := engine.Move(from, to, movesWhite); err != nil {
		t.Errorf("%s: %s -> %s moving white %v. Expected err = nil, got %v", testName, from, to, movesWhite, err)
	}
}
