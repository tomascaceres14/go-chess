package engine

type pawn struct {
	*basePiece
	direction int
}

func newPawn(pos position, p *player) *pawn {
	white := p.isWhite

	dir := 1
	if !white {
		dir = -1
	}

	pawn := &pawn{
		basePiece: newBasePiece(white, 1, pos, nil),
		direction: dir,
	}

	p.pieces = append(p.pieces, pawn)

	return pawn
}

func (p *pawn) visibleSquares(b *board) map[position]bool {

	positions := map[position]bool{}
	front := position{Row: p.pos.Row + 1*p.direction, Col: p.pos.Col}
	diag1 := position{Row: p.pos.Row + 1*p.direction, Col: p.pos.Col + 1}
	diag2 := position{Row: p.pos.Row + 1*p.direction, Col: p.pos.Col - 1}

	if diag1.inBounds() {
		positions[diag1] = true
	}

	if diag2.inBounds() {
		positions[diag2] = true
	}

	positions[front] = true

	if !p.moved {
		front.Row += 1 * p.direction
		positions[front] = true
	}

	return positions
}

func (p *pawn) legalMoves(b *board) map[position]bool {

	positions := p.visibleSquares(b)
	legalMoves := map[position]bool{}

	for pos := range positions {

		piece, occupied := b.getPiece(pos)

		// move front
		if pos.Col == p.pos.Col {
			if !occupied {
				legalMoves[pos] = true
			}
			continue
		}

		// capture diagonal
		if occupied {
			if piece.isWhite() != p.white {
				legalMoves[pos] = true
			}
			continue
		}

		// TODO: en passant
	}

	return legalMoves
}

func (p *pawn) getPosition() position {
	return p.pos
}

func (p *pawn) setPosition(pos position) {
	p.moved = true
	p.pos = pos
}

func (p *pawn) isWhite() bool {
	return p.white
}

func (p *pawn) String() string {
	piece := "P"

	if !p.white {
		piece = "p"
	}

	return piece
}

func (p *pawn) clone() movable {
	return &pawn{basePiece: p.basePiece.cloneBase()}
}

func (p *pawn) getType() pieceType {
	return pawnType
}
