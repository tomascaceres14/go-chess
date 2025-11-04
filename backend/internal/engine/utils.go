package engine

import (
	"strconv"
)

// Board columns
const cols = "abcdefgh"

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

func IsMoveSafeToKing(piece Movable, to Position, g *Game) bool {

	// Clone board
	board := g.board.Clone()

	// Simulate movement on cloned board
	board.MovePieceSim(piece.GetPosition(), to)

	// Calculate threats of opponent but using the individual pieces of the cloned board, not the player.Pieces slice.
	// The latter would require to also deep copy the whole game.
	opponentColor := !piece.IsWhite()
	threats := AttackedByColor(board, opponentColor)

	// Find king on cloned board
	kingPos, ok := FindKingPos(board, piece.IsWhite())
	if !ok {
		return false
	}

	// If king was moved, then 'to' is its safe square
	if piece == g.GetPlayer(piece.IsWhite()).King {
		kingPos = to
	}

	return !threats[kingPos]
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

func ColorToString(isWhite bool) string {
	if isWhite {
		return "Whites"
	}
	return "Blacks"
}

// Used for simulating movements
// Calculate AttaquedSquares but using the board, not the player.Pieces slice
func AttackedByColor(b *Board, white bool) map[Position]bool {
	threats := make(map[Position]bool)
	for i := range 8 {
		for j := range 8 {
			p := (*b.grid)[i][j]
			if p != nil && p.IsWhite() == white {
				for sq := range p.VisibleSquares(b) {
					threats[sq] = true
				}
			}
		}
	}
	return threats
}

// Used for simulating movements
// Find king by color using board
func FindKingPos(b *Board, white bool) (Position, bool) {
	for i := range 8 {
		for j := range 8 {
			p := (*b.grid)[i][j]
			if p != nil && p.IsWhite() == white {
				if _, ok := p.(*King); ok {
					return p.GetPosition(), true
				}
			}
		}
	}
	return Position{}, false
}

func GETFENPosition(g *Game) string {
	FENString := ""
	grid := g.board.grid
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

		if i == len(g.board.grid)-1 {
			endLine = ""
		}

		FENString += endLine
	}
	return FENString
}

func GetFENCastling(g *Game) string {
	FENString := ""

	wKing := g.PWhite.King
	bKing := g.PBlack.King

	if wKing.ShortCastlingOpt {
		FENString += "K"
	}

	if wKing.LongCastlingOpt {
		FENString += "Q"
	}

	if bKing.ShortCastlingOpt {
		FENString += "k"
	}

	if bKing.LongCastlingOpt {
		FENString += "q"
	}

	if FENString == "" {
		FENString = "-"
	}

	return FENString
}

func GetFENString(g *Game) string {

	FENString := GETFENPosition(g)

	// Define turn of player
	turn := "w"
	if !g.WhiteTurn {
		turn = "b"
	}

	FENString += " " + turn + " "

	FENString += GetFENCastling(g) + " "

	return FENString
}
