package gochess

type move struct {
	PieceCopy         movable
	From, To          position
	Capture           movable
	AlgebraicNotation string
	IsCheck           bool
}

func newMove(piece, capture movable, from, to position, isCheck bool, castleDir int) move {
	algebraicNotation := piece.getAlgebraicString()

	takes := "x"
	check := ""

	if capture == nil {
		takes = ""
	} else if piece.getType() == pawnType {
		takes = from.getCol() + takes
	}

	if isCheck {
		check = "+"
	}

	switch castleDir {
	case -1:
		algebraicNotation = algebraicNotation + takes + to.String() + check
	case 0:
		algebraicNotation = "0-0"
	case 1:
		algebraicNotation = "0-0-0"
	}

	return move{
		PieceCopy:         piece,
		Capture:           capture,
		From:              from,
		To:                to,
		AlgebraicNotation: algebraicNotation,
	}
}

func (m move) String() string {
	return m.AlgebraicNotation
}
