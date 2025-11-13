package gochess

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

	rank := pos.getRank()

	pawn.moved = !(white && rank == 2)
	pawn.moved = !(!white && rank == 7)

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

			// if pawn is jumping one square
			if pos.row == p.pos.row+2*p.direction && !occupied && !b.IsOccupied(pos) {
				legalMoves[pos] = true
				continue
			}

		} else {
			// capture diagonal
			regularCapture := occupied && piece.isWhite() != p.white
			legalMoves[pos] = regularCapture

			// check if is in rank of en passant
			if p.pos.getRank() != 5 && p.pos.getRank() != 4 {
				continue
			}

			// create en passant position
			enPassant := pos
			enPassant.row = enPassant.row - 1*p.direction

			// get piece at en passant square. Continue if empty
			enPassantMovable, enPassantOcc := b.getPiece(enPassant)
			if !enPassantOcc {
				continue
			}

			// Cast to pawn. Continue if not pawn
			enPassantPawn, ok := castPawn(enPassantMovable)
			if !ok {
				continue
			}

			// Check pawn has jumped and its not white
			enPassantCapture := enPassantPawn.jumped && enPassantPawn.white != p.white
			legalMoves[pos] = regularCapture || enPassantCapture
		}
	}
	return legalMoves
}

func (p *pawn) move(to position, game *game) movable {
	from := p.pos
	board := game.gameBoard
	player := game.GetPlayer(p.white)
	capture := board.movePiece(p, to)
	p.setPosition(to)
	p.setMoved(true)

	// Promotion
	if to.row == 0 || to.row == 7 {
		queen := newQueen(to, player)
		board.insertPiece(queen)
		queen.setMoved(true)
		return capture
	}

	// En passant
	if to.col != from.col {
		capturedPos := position{row: from.row, col: to.col}
		capture, _ = board.getPiece(capturedPos)
		board.clearSquare(capturedPos)
		return capture
	}

	diff := from.row - to.row
	pawnJumped := diff == 2 || diff == -2

	if pawnJumped {
		p.jumped = pawnJumped
		player.pawnJumped = p
	}

	return capture
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
