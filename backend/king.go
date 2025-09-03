package main

type King struct {
	*BasePiece
}

func NewKing(pos Position, p *Player) *King {
	white := p.White
	directions := []Direction{
		// left
		{-1, 0},
		// right
		{1, 0},
		// up
		{0, 1},
		// down
		{0, -1},
		// top right
		{1, 1},
		// top left
		{-1, 1},
		// bottom right
		{1, -1},
		// bottom left
		{-1, -1},
	}

	king := &King{
		BasePiece: NewBasePiece(white, 0, pos, directions),
	}

	p.Pieces = append(p.Pieces, king)

	return king
}

func (k *King) VisibleSquares(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range k.Directions {
		pos := Position{Row: k.Pos.Row + v.dx, Col: k.Pos.Col + v.dy}

		if !pos.InBounds() {
			continue
		}

		pieceAt, occupied := b.GetPiece(pos)

		if !occupied || pieceAt.IsWhite() != k.White {
			positions[pos] = true
			continue
		}
	}

	return positions
}

func (k *King) LegalMoves(b *Board) map[Position]bool {
	legalMoves := k.VisibleSquares(b)

	if k.hasMoved {
		return legalMoves
	}

	shortCastlePos := Position{Row: k.Pos.Row, Col: 7}
	longCastlePos := Position{Row: k.Pos.Row, Col: 0}

	shortRook, occ := b.GetPiece(shortCastlePos)
	if occ && !shortRook.HasMoved() && b.IsRowPathClear(k.Pos, shortCastlePos) {
		legalMoves[Position{Row: k.Pos.Row, Col: 6}] = true
	}

	longRook, occ := b.GetPiece(longCastlePos)
	if occ && !longRook.HasMoved() && b.IsRowPathClear(k.Pos, longCastlePos) {
		legalMoves[Position{Row: k.Pos.Row, Col: 2}] = true
	}

	return legalMoves
}

func (k *King) GetPosition() Position {
	return k.Pos
}

func (k *King) SetPosition(pos Position) {
	k.hasMoved = true
	k.Pos = pos
}

func (k *King) IsWhite() bool {
	return k.White
}

func (k *King) String() string {
	piece := "♔"

	if !k.White {
		piece = "♚"
	}

	return piece
}

func (k *King) Clone() Movable {
	return &King{BasePiece: k.BasePiece.CloneBase()}
}

func (k *King) GetType() PieceType {
	return KingType
}
