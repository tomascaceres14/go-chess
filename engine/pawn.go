package gochess

import "fmt"

type pawn struct {
	*basePiece
	direction int
}

func newPawn(pos position, p *player) *pawn {

	white := p.isWhite

	const WhiteInitialRank = 2
	const BlackInitialRank = 7

	dir := 1
	if !white {
		dir = -1
	}

	pawn := &pawn{
		basePiece: newBasePiece(white, 1, pos, nil),
		direction: dir,
	}

	rank := pos.getRank()

	if white {
		pawn.moved = rank != WhiteInitialRank
	} else {
		pawn.moved = rank != BlackInitialRank
	}

	p.pieces = append(p.pieces, pawn)

	return pawn
}

func (p *pawn) visibleSquares(b *board) map[position]bool {

	positions := map[position]bool{}

	front1 := position{row: p.pos.row + 1*p.direction, col: p.pos.col}
	front2 := position{row: p.pos.row + 2*p.direction, col: p.pos.col}
	diag1 := position{row: p.pos.row + 1*p.direction, col: p.pos.col + 1}
	diag2 := position{row: p.pos.row + 1*p.direction, col: p.pos.col - 1}

	if diag1.inBounds() {
		positions[diag1] = true
	}

	if diag2.inBounds() {
		positions[diag2] = true
	}
	if front1.inBounds() {
		positions[front1] = true
	}

	if !p.moved && front2.inBounds() {
		positions[front2] = true
	}

	return positions
}

func (p *pawn) legalMoves(b *board) map[position]bool {

	positions := p.visibleSquares(b)
	legalMoves := map[position]bool{}
	for pos := range positions {

		piece, occupied := b.getPiece(pos)

		if pos.col == p.pos.col {

			// if pawn moving one square up
			if pos.row == p.pos.row+1*p.direction && !occupied {
				legalMoves[pos] = true
				continue
			}

			// if pawn is jumping
			if pos.row == p.pos.row+2*p.direction && !occupied && !b.IsOccupied(pos) {
				legalMoves[pos] = true
				continue
			}

		} else {
			// capture diagonal
			regularCapture := occupied && piece.isWhite() != p.white
			EPTarget := b.enPassantTarget
			enPassantCapture := false
			if EPTarget != nil {
				enPassantCapture = b.enPassantTarget.equals(pos)
			}
			legalMoves[pos] = regularCapture || enPassantCapture
		}
	}
	return legalMoves
}

func (p *pawn) move(to position, game *game) movable {

	from := p.pos
	board := game.gameBoard

	capture := board.movePiece(p, to)
	p.setPosition(to)

	// calculate if its trying to jump
	if !p.moved {

		diff := from.row - to.row
		if diff == 2 || diff == -2 {
			mid := position{row: to.row - 1*p.direction, col: from.col}
			board.enPassantTarget = &mid
		}

		p.setMoved(true)
		return capture
	}

	rank := to.getRank()

	// Promotion
	if rank == 1 || rank == 8 {
		queen := newQueen(to, game.GetPlayer(p.white))
		board.insertPiece(queen)
		queen.setMoved(true)
		return capture
	}

	// En passant
	EPTarget := board.enPassantTarget
	if EPTarget != nil && EPTarget.equals(to) {
		fmt.Println("capturing enpassant from", from, "to", to)
		capturedPos := position{row: from.row - 1*p.direction, col: to.col}
		capture, _ = board.getPiece(capturedPos)
		board.clearSquare(capturedPos)
		board.enPassantTarget = nil
		return capture
	}

	return capture
}

func (p *pawn) getPosition() position {
	return p.pos
}

func (p *pawn) setPosition(pos position) {
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
	return &pawn{basePiece: p.basePiece.cloneBase(), direction: p.direction}
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
