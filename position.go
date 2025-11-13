package gochess

import (
	"fmt"
	"strconv"
	"strings"
)

type position struct {
	row, col int
}

type direction struct {
	dx, dy int
}

func (p *position) isValid() bool {
	return p.row != -1 && p.col != -1
}

func pos(pos string) position {
	nilPos := position{-1, -1}

	if len(pos) != 2 {
		return nilPos
	}

	col := strings.ToLower(string(pos[0]))
	row, err := strconv.Atoi(string(pos[1]))
	if err != nil {
		return nilPos
	}

	if (row < 1 || row > 8) || (col[0] < 'a' || col[0] > 'h') {
		return nilPos
	}

	newPos := position{
		row: row - 1,
		col: int(col[0] - 'a'),
	}
	return newPos
}

func (p position) inBounds() bool {
	return (0 <= p.col && p.col <= 7) && (0 <= p.row && p.row <= 7)
}

func (p position) getFile() string {
	return string(cols[p.col])
}

func (p position) getRank() int {
	return p.row + 1
}

func (p position) String() string {
	return fmt.Sprintf("%v%v", string(cols[p.col]), p.row+1)
}
