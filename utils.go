package gochess

import (
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
	next := position{row: pos.row + dx, col: pos.col + dy}
	castRay(next, dx, dy, b, white, positions)
}

func isMoveSafeToKing(piece movable, to position, board *board) bool {

	// Clone board
	boardSim := board.clone()

	// Simulate movement on cloned board
	boardSim.MovePieceSim(piece.getPosition(), to)

	var kingPos position
	if piece.getType() == kingType {
		kingPos = to
	} else {
		kingPos = boardSim.findKingPos(piece.isWhite())
	}

	return !boardSim.isKingInCheck(kingPos, piece.isWhite())
}

// Parse col from matrix index to board column letter
func getColLetter(col int) string {
	return string(cols[col])
}

func getFENPosition(g *game) string {
	FENString := ""
	grid := g.gameBoard.grid

	// in engine, grid starts with a1 being [0][0] so to generate FEN pos its necessary to do reverse traversal.
	for i := len(grid) - 1; i >= 0; i-- {
		row := grid[i]
		emptySquares := 0

		if len(row) == 0 {
			emptySquares = 8
		}

		for j := range len(row) {
			v := row[j]
			if v != nil {
				if emptySquares > 0 {
					FENString += strconv.Itoa(emptySquares)
					emptySquares = 0
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

		if i == 0 {
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

	return FENString + " "
}

func getFENEnPassant(g *game) string {
	FENString := "- "
	turn := g.getCurrentTurn()
	player := g.GetPlayer(!turn)

	pawn := player.pawnJumped

	if pawn != nil {
		prevSquare := position{col: pawn.pos.col, row: pawn.pos.row - 1*pawn.direction}
		FENString = prevSquare.String() + " "
	}

	return FENString
}

func colorToString(color bool) string {
	colorStr := "white"

	if !color {
		colorStr = "black"
	}

	return colorStr
}
