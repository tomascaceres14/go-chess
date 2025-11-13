package gochess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type player struct {
	name               string
	isWhite, isChecked bool
	points             int
	king               *king
	pawnJumped         *pawn
	pieces             []movable
	threats            map[position]bool
}

func newPlayer(name string, isWhite bool) (*player, error) {
	name = strings.TrimSpace(name)

	if len(name) <= 0 {
		return nil, errors.New("Player name should include at least one character.")
	}
	return &player{
		name:    name,
		isWhite: isWhite,
		points:  0,
	}, nil
}

func newPlayerWhite(name string) (*player, error) {
	p, err := newPlayer(name, true)
	return p, err
}

func newPlayerBlack(name string) (*player, error) {
	p, err := newPlayer(name, false)
	return p, err
}

func (p *player) legalMoves(g *game) map[position]bool {
	moves := make(map[position]bool)

	pieces := p.pieces
	for _, piece := range pieces {
		for pos := range piece.legalMoves(g.gameBoard) {
			if isMoveSafeToKing(piece, pos, g.gameBoard) {
				moves[pos] = true
			}
		}
	}

	return moves
}

func (p *player) hasLegalMoves(g *game) bool {
	return len(p.legalMoves(g)) > 0
}

func (p *player) removeJumpFromPawn() {
	p.pawnJumped.jumped = false
	p.pawnJumped = nil
}

func (p *player) getKing() *king {
	return p.king
}

func (p *player) deletePiece(piece movable) {
	pieces := p.pieces
	for i, v := range pieces {
		if v == piece {
			p.pieces = append(pieces[0:i], pieces[i+1:]...)
			return
		}
	}
}

func (p *player) incrementPoints(inc int) {
	p.points += inc
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
