package main

import (
	"fmt"
)

// func randomBool() bool {
// 	return rand.Intn(2) == 0
// }

const cols = "ABCDEFGH"

func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

func CastRay(pos Position, dx, dy int, b *Board, white bool, positions map[Position]bool) {

	if !pos.InBounds() {
		return
	}

	piece, occupied := b.GetPiece(pos)
	if occupied {
		if piece.IsWhite() != white {
			positions[pos] = true
		}
		return
	}

	positions[pos] = true
	next := Position{Row: pos.Row + dx, Col: pos.Col + dy}
	CastRay(next, dx, dy, b, white, positions)
}

func GetCol(col int) string {
	return string(cols[col])
}

func GetRow(row int) int {
	return row + 1
}
