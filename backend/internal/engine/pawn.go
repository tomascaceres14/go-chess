package engine

import "fmt"

type pawn struct {
	*basePiece
	direction int
	jumped    bool
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
	front1 := position{Row: p.pos.Row + 1*p.direction, Col: p.pos.Col}
	front2 := position{Row: p.pos.Row + 2*p.direction, Col: p.pos.Col}
	diag1 := position{Row: p.pos.Row + 1*p.direction, Col: p.pos.Col + 1}
	diag2 := position{Row: p.pos.Row + 1*p.direction, Col: p.pos.Col - 1}

	if diag1.inBounds() {
		positions[diag1] = true
	}

	if diag2.inBounds() {
		positions[diag2] = true
	}

	positions[front1] = true

	if !p.moved {
		positions[front2] = true
	}

	return positions
}

func (p *pawn) legalMoves(b *board) map[position]bool {

	positions := p.visibleSquares(b)
	legalMoves := map[position]bool{}

	for pos := range positions {

		piece, occupied := b.getPiece(pos)

		if pos.Col == p.pos.Col {

			// if pawn moving one square up
			if pos.Row == p.pos.Row+1*p.direction && !occupied {
				legalMoves[pos] = true
				continue
			}

			// if pawn is jumping one square
			if pos.Row == p.pos.Row+2*p.direction && !occupied && !b.IsOccupied(pos) {
				legalMoves[pos] = true
				continue
			}

		} else {
			// capture diagonal
			legalMoves[pos] = occupied && piece.isWhite() != p.white
		}

	}

	fmt.Println("Legal moves for pawn at: ", p.pos, legalMoves)

	// en passant. Not checking if pawn not in 6th or 3rd rank.
	if p.pos.getRow() != 6 || p.pos.getRow() != 3 {
		return legalMoves
	}
	println("pawn maybe can enpassant")
	// Define left square
	leftPos := position{Row: p.pos.Row, Col: p.pos.Col - 1}
	// Get piece
	leftMovable, occ := b.getPiece(leftPos)
	// Cast to pawn
	leftPawn, ok := castPawn(leftMovable)
	// Verify left square is occupied, piece is pawn, its from opposite color and has jumped
	legalMoves[leftPos] = occ && ok && leftPawn.isWhite() != p.white && leftPawn.jumped

	// Same for right square
	rightPos := position{Row: p.pos.Row, Col: p.pos.Col - 1}
	rightMovable, occ := b.getPiece(rightPos)
	rightPawn, ok := castPawn(rightMovable)
	legalMoves[rightPos] = occ && ok && rightPawn.isWhite() != p.white && rightPawn.jumped

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

func castPawn(m movable) (*pawn, bool) {

	if m == nil || m.getType() != pawnType {
		return &pawn{}, false
	}

	pawn, _ := m.(*pawn)

	return pawn, true
}
