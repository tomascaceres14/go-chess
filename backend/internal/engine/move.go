package engine

type move struct {
	PieceCopy         movable
	From, To          position
	Capture           movable
	AlgebraicNotation string
	IsCheck           bool
}

func newMove(piece, capture movable, from, to position, isCheck bool) move {
	algebraicNotation := piece.getAlgebraicString()

	takes := "x"
	check := ""

	if capture == nil {
		takes = ""
	} else {
		if _, ok := piece.(*pawn); ok {
			takes = from.getCol() + takes
		}
	}

	if isCheck {
		check = "+"
	}

	algebraicNotation = algebraicNotation + takes + to.String() + check

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
