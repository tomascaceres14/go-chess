package main

import (
	"fmt"
	"strings"
)

type Position struct {
	Row, Col int
}

func Pos(col string, row int) Position {

	if row < 1 || row > 8 {
		return Position{-1, -1}
	}

	col = strings.ToLower(col)
	if len(col) != 1 || col[0] < 'a' || col[0] > 'h' {
		return Position{-1, -1}
	}

	return Position{
		Row: row - 1,
		Col: int(col[0] - 'a'),
	}
}

func (p Position) Equals(other Position) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func (p Position) InBounds() bool {
	return (0 <= p.Col && p.Col <= 7) && (0 <= p.Row && p.Row <= 7)
}

func (p Position) String() string {
	return fmt.Sprintf("{%v%v}", string(cols[p.Col]), p.Row+1)
}

func ContainsPosition(list []Position, pos Position) bool {
	for _, p := range list {
		if p.Equals(pos) {
			return true
		}
	}

	return false
}
