package gochess

type move struct {
	piece             movable
	from, to          position
	capture           movable
	algebraicNotation string
	isCheck           bool
	color             bool
	castleDir         int
}

func (m move) String() string {
	return m.getAlgebraicString()
}

func (m *move) getAlgebraicString() string {

	if m.piece == nil {
		return "-"
	}

	if m.algebraicNotation != "" {
		return m.algebraicNotation
	}

	if m.piece.getType() == kingType {
		k, _ := castKing(m.piece)
		m.castleDir = k.castleDir
		switch k.castleDir {
		case 0:
			return "0-0"
		case 1:
			return "0-0-0"
		}
	}

	algebraicNotation := m.piece.getAlgebraicString()

	takes := "x"
	check := ""

	if m.capture == nil {
		takes = ""
	} else if m.piece.getType() == pawnType {
		takes = m.from.getFile() + takes
	}

	if m.isCheck {
		check = "+"
	}

	m.algebraicNotation = algebraicNotation + takes + m.to.String() + check

	return m.algebraicNotation

}

func (m *move) getPiece() movable {
	return m.piece
}
