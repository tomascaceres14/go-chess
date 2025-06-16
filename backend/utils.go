package main

import (
	"fmt"
)

const cols = "ABCDEFGH"

// Error printing for debugging
func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

// Recursive function to cast a ray and check for collisions in direction vector {dx, dy}.
// Returns map of possible positions
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

// Parse col from matrix index to board column letter
func GetCol(col int) string {
	return string(cols[col])
}

// Parse col from matrix index to board column letter
func GetRow(row int) int {
	return row + 1
}

// Removes a piece from a list of pieces
func DeletePiece(list []Movable, piece Movable) []Movable {
	for i, v := range list {
		if v == piece {
			return append(list[0:i], list[i+1:]...)
		}
	}

	return list
}

// func randomBool() bool {
// 	return rand.Intn(2) == 0
// }
