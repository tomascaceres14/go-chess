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

func (p *King) AttackedSquares(b *Board) map[Position]bool {
	positions := map[Position]bool{}

	for _, v := range p.Directions {
		pos := Position{Row: p.Pos.Row + v.dx, Col: p.Pos.Col + v.dy}

		if !pos.InBounds() {
			continue
		}

		pieceAt, occupied := b.GetPiece(pos)

		if !occupied || pieceAt.IsWhite() != p.White {
			positions[pos] = true
			continue
		}
	}

	return positions
}

func (p *King) LegalMoves(b *Board) map[Position]bool {
	legalMoves := p.LegalMovesDefault(b)

	if p.hasMoved {
		return legalMoves
	}

	shortCastlePos := Position{Row: p.Pos.Row, Col: 7}
	longCastlePos := Position{Row: p.Pos.Row, Col: 0}

	shortRook, occ := b.GetPiece(shortCastlePos)
	if occ && !shortRook.HasMoved() && b.IsRowPathClear(p.Pos, shortCastlePos) {
		legalMoves[Position{Row: p.Pos.Row, Col: 6}] = true
	}

	longRook, occ := b.GetPiece(longCastlePos)
	if occ && !longRook.HasMoved() && b.IsRowPathClear(p.Pos, longCastlePos) {
		legalMoves[Position{Row: p.Pos.Row, Col: 2}] = true
	}

	return legalMoves
}

func (p *King) GetPosition() Position {
	return p.Pos
}

func (p *King) SetPosition(pos Position) {
	p.hasMoved = true
	p.Pos = pos
}

func (p *King) IsWhite() bool {
	return p.White
}

func (p *King) String() string {
	color := "w"

	if !p.White {
		color = "b"
	}

	return "K" + color
}

func (p *King) Clone() Movable {
	return &King{BasePiece: p.BasePiece.CloneBase()}
}
