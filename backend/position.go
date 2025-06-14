package main

import (
	"fmt"
	"strings"
)

type Position struct {
	Row, Col int
}

func Pos(row int, col string) (Position, error) {
	if row < 1 || row > 8 {
		return Position{}, fmt.Errorf("row fuera de rango: %d", row)
	}

	col = strings.ToLower(col)
	if len(col) != 1 || col[0] < 'a' || col[0] > 'h' {
		return Position{}, fmt.Errorf("columna inválida: %s", col)
	}

	return Position{
		Row: row - 1,
		Col: int(col[0] - 'a'),
	}, nil
}

func (p Position) Equals(other Position) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func ContainsPosition(list []Position, pos Position) bool {
	for _, p := range list {
		if p.Equals(pos) {
			return true
		}
	}

	return false
}
