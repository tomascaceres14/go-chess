package engine

type Move struct {
	PieceCopy         Movable
	From, To          Position
	Capture           Movable
	AlgebraicNotation string
	IsCheck           bool
}

func NewMove(piece, capture Movable, from, to Position, isCheck bool) Move {
	algebraicNotation := piece.GetAlgebraicString()

	takes := "x"
	check := ""

	if capture == nil {
		takes = ""
	} else {
		if _, ok := piece.(*Pawn); ok {
			takes = from.GetCol() + takes
		}
	}

	if isCheck {
		check = "+"
	}

	algebraicNotation = algebraicNotation + takes + to.String() + check

	return Move{
		PieceCopy:         piece,
		Capture:           capture,
		From:              from,
		To:                to,
		AlgebraicNotation: algebraicNotation,
	}
}

func (m Move) String() string {
	return m.AlgebraicNotation
}
