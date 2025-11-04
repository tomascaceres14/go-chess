package engine

import (
	"fmt"
	"strconv"
)

type player struct {
	name      string
	isWhite   bool
	points    int
	king      *king
	isChecked bool
	pieces    []movable
	threats   map[position]bool
}

func newPlayer(name string, isWhite bool) *player {
	return &player{
		name:    name,
		isWhite: isWhite,
		points:  0,
		threats: map[position]bool{},
	}
}

func (p *player) attackedSquares(b *board) map[position]bool {
	threats := make(map[position]bool)

	pieces := p.pieces
	for _, v := range pieces {
		for k := range v.visibleSquares(b) {
			threats[k] = true
		}
	}

	return threats
}

func (p *player) legalMoves(g *game) map[position]bool {
	moves := make(map[position]bool)

	pieces := p.pieces
	for _, piece := range pieces {
		for pos := range piece.legalMoves(g.gameBoard) {
			if isMoveSafeToKing(piece, pos, g) {
				moves[pos] = true
			}
		}
	}

	return moves
}

func (p *player) hasLegalMoves(g *game) bool {
	return len(p.legalMoves(g)) > 0
}

func (p *player) String() string {
	col := "WHITE"

	if !p.isWhite {
		col = "BLACK"
	}

	result := fmt.Sprintf("%s plays %s:\n\tPoints: %v\n\tChecked: %s\n\tPieces on board: %v",
		p.name,
		col,
		strconv.Itoa(p.points),
		strconv.FormatBool(p.isChecked),
		p.pieces)

	return result
}
