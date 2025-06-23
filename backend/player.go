package main

import (
	"fmt"
	"strconv"
)

type Player struct {
	Name    string
	White   bool
	Points  int
	King    *King
	Checked bool
	Pieces  []Movable
	Threats map[Position]bool
}

func NewPlayer(name string, isWhite bool) *Player {
	return &Player{
		Name:    name,
		White:   isWhite,
		Points:  0,
		Threats: map[Position]bool{},
	}
}

func (p *Player) AttackedSquares(b *Board) map[Position]bool {
	threats := make(map[Position]bool)

	pieces := p.Pieces
	for _, v := range pieces {
		for k := range v.AttackedSquares(b) {
			threats[k] = true
		}
	}

	return threats
}

func (p *Player) String() string {
	col := "WHITE"
	if !p.White {
		col = "BLACK"
	}
	result := fmt.Sprintf("%s plays %s:\n\tPoints: %v\n\tChecked: %s\n\tPieces on board: %v",
		p.Name,
		col,
		strconv.Itoa(p.Points),
		strconv.FormatBool(p.Checked),
		p.Pieces)

	return result
}
