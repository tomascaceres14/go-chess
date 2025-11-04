package engine

import (
	"fmt"
	"strconv"
)

// Board columns
const cols = "abcdefgh"

// Recursive function to cast a ray and check for collisions in direction vector {dx, dy}.
// Returns map of possible positions
func castRay(pos position, dx, dy int, b *board, white bool, positions map[position]bool) {

	if !pos.inBounds() {
		return
	}

	piece, occupied := b.getPiece(pos)
	if occupied {
		if piece.isWhite() != white {
			positions[pos] = true
		}
		return
	}

	positions[pos] = true
	next := position{Row: pos.Row + dx, Col: pos.Col + dy}
	castRay(next, dx, dy, b, white, positions)
}

// Error printing for debugging
func printError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

func isMoveSafeToKing(piece movable, to position, g *game) bool {

	// Clone board
	board := g.gameBoard.Clone()

	// Simulate movement on cloned board
	board.MovePieceSim(piece.getPosition(), to)

	// Calculate threats of opponent but using the individual pieces of the cloned board, not the player.Pieces slice.
	// The latter would require to also deep copy the whole game.
	opponentColor := !piece.isWhite()
	threats := attackedByColor(board, opponentColor)

	// Find king on cloned board
	kingPos, ok := findKingPos(board, piece.isWhite())
	if !ok {
		return false
	}

	// If king was moved, then 'to' is its safe square
	if piece == g.GetPlayer(piece.isWhite()).king {
		kingPos = to
	}

	return !threats[kingPos]
}

// Parse col from matrix index to board column letter
func getCol(col int) string {
	return string(cols[col])
}

// Parse col from matrix index to board column letter
func getRow(row int) int {
	return row + 1
}

// Removes a piece from a list of pieces
func deletePiece(list []movable, piece movable) []movable {
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

func colorToString(isWhite bool) string {
	if isWhite {
		return "Whites"
	}
	return "Blacks"
}

// Used for simulating movements
// Calculate AttaquedSquares but using the board, not the player.Pieces slice
func attackedByColor(b *board, white bool) map[position]bool {
	threats := make(map[position]bool)
	for i := range 8 {
		for j := range 8 {
			p := (*b.grid)[i][j]
			if p != nil && p.isWhite() == white {
				for sq := range p.visibleSquares(b) {
					threats[sq] = true
				}
			}
		}
	}
	return threats
}

// Used for simulating movements
// Find king by color using board
func findKingPos(b *board, white bool) (position, bool) {
	for i := range 8 {
		for j := range 8 {
			p := (*b.grid)[i][j]
			if p != nil && p.isWhite() == white {
				if _, ok := p.(*king); ok {
					return p.getPosition(), true
				}
			}
		}
	}
	return position{}, false
}

func getFENPosition(g *game) string {
	FENString := ""
	grid := g.gameBoard.grid
	for i := len(grid) - 1; i >= 0; i-- {
		row := grid[i]
		emptySquares := 0

		if len(row) == 0 {
			emptySquares = 8
		}

		for j := len(row) - 1; j >= 0; j-- {
			v := row[j]
			if v != nil {
				if emptySquares > 0 {
					FENString += strconv.Itoa(emptySquares)
				}
				FENString += v.String()
			} else {
				emptySquares++
			}
		}

		if emptySquares > 0 {
			FENString += strconv.Itoa(emptySquares)
		}

		endLine := "/"

		if i == len(g.gameBoard.grid)-1 {
			endLine = ""
		}

		FENString += endLine
	}
	return FENString
}

func getFENCastling(g *game) string {
	FENString := ""

	wKing := g.pWhite.king
	bKing := g.pBlack.king

	if wKing.shortCastlingOpt {
		FENString += "K"
	}

	if wKing.longCastlingOpt {
		FENString += "Q"
	}

	if bKing.shortCastlingOpt {
		FENString += "k"
	}

	if bKing.longCastlingOpt {
		FENString += "q"
	}

	if FENString == "" {
		FENString = "-"
	}

	return FENString
}

func (g *game) GetFENString() string {

	FENString := getFENPosition(g)

	// Define turn of player
	turn := "w"
	if !g.WhiteTurn {
		turn = "b"
	}

	FENString += " " + turn + " "

	FENString += getFENCastling(g) + " "

	return FENString
}
