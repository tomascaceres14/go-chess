package gochess

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
	"unicode"
)

type gameStatus int

const (
	aborted gameStatus = iota
	playing
	whiteWins
	blackWins
	draw
)

type game struct {
	id             string
	gameBoard      *board
	pWhite, pBlack *player
	WhiteTurn      bool
	halfmoveClock  int
	moveHistory    []move
	status         gameStatus
}

// Generates a new board with classic chess configuration
func NewDefaultGame(whiteName, blackName string) (*game, error) {

	fmt.Println("generating new board...")

	gameBoard := [8][8]movable{}

	pWhite, err := newPlayerWhite(whiteName)
	if err != nil {
		return nil, err
	}
	pBlack, err := newPlayerBlack(blackName)
	if err != nil {
		return nil, err
	}

	// Generate ID for match
	now := time.Now()
	timestamp := now.Format("20060201150405")
	id := whiteName + "_" + blackName + "_" + timestamp

	game := &game{
		id:            id,
		gameBoard:     &board{grid: &gameBoard},
		pWhite:        pWhite,
		pBlack:        pBlack,
		WhiteTurn:     true,
		moveHistory:   []move{},
		status:        playing,
		halfmoveClock: 0,
	}

	// Pawns
	for i := range 8 {
		game.gameBoard.insertPiece(newPawn(pos(getColLetter(i)+"7"), pBlack)) // black
		game.gameBoard.insertPiece(newPawn(pos(getColLetter(i)+"2"), pWhite)) // white
	}

	// Rooks
	game.gameBoard.insertPiece(newRook(pos("a8"), pBlack)) // black
	game.gameBoard.insertPiece(newRook(pos("h8"), pBlack)) // black
	game.gameBoard.insertPiece(newRook(pos("a1"), pWhite)) // white
	game.gameBoard.insertPiece(newRook(pos("h1"), pWhite)) // white

	// Knights
	game.gameBoard.insertPiece(newKnight(pos("b8"), pBlack)) // black
	game.gameBoard.insertPiece(newKnight(pos("g8"), pBlack)) // black
	game.gameBoard.insertPiece(newKnight(pos("b1"), pWhite)) // white
	game.gameBoard.insertPiece(newKnight(pos("g1"), pWhite)) // white

	// Bishops
	game.gameBoard.insertPiece(newBishop(pos("c8"), pBlack)) // black
	game.gameBoard.insertPiece(newBishop(pos("f8"), pBlack)) // black
	game.gameBoard.insertPiece(newBishop(pos("c1"), pWhite)) // white
	game.gameBoard.insertPiece(newBishop(pos("f1"), pWhite)) // white

	// Queens
	game.gameBoard.insertPiece(newQueen(pos("d8"), pBlack)) // black
	game.gameBoard.insertPiece(newQueen(pos("d1"), pWhite)) // white

	// Kings
	bKing := newKing(pos("e8"), pBlack)
	pBlack.king = bKing
	game.gameBoard.insertPiece(bKing)

	wKing := newKing(pos("e1"), pWhite)
	pWhite.king = wKing
	game.gameBoard.insertPiece(wKing)

	fmt.Println(game.gameBoard)

	return game, nil
}

// Obtains piece at given position if player is owner of piece
func (g *game) getPlayerPiece(pos position, pColor bool) (movable, error) {

	piece, ok := g.gameBoard.getPiece(pos)
	if !ok {
		return nil, fmt.Errorf("Position %v is empty.", pos)
	}

	// Validate piece belonging to player
	if piece.isWhite() != pColor {
		return nil, fmt.Errorf("Not your piece, %s.", colorToString(pColor))
	}

	return piece, nil
}

// Moves piece in position `from` to position `to` if player is owner of piece
func (game *game) makeMove(from, to position, pColor bool) error {

	// Check turn to play
	if game.WhiteTurn != pColor {
		return fmt.Errorf("Not your turn, %s.", colorToString(pColor))
	}

	// Check if positions are in bounds
	if !from.inBounds() || !to.inBounds() {
		return errors.New("Position out of bounds.")
	}

	// Obtain player and opponent
	player := game.GetPlayer(pColor)
	opponent := game.GetPlayerOpponent(player.isWhite)

	// Obtains piece to move
	piece, err := game.getPlayerPiece(from, player.isWhite)
	if err != nil {
		return err
	}

	// Check if piece can move to desired position or if is trying to move in-place
	legalMoves := piece.legalMoves(game.gameBoard)
	if !legalMoves[to] || piece.getPosition() == to {
		return fmt.Errorf("%s can't move from %s to %s.", piece.String(), from, to)
	}

	// Check if the move leaves king vulnerable
	if !isMoveSafeToKing(piece, to, game) {
		return fmt.Errorf("%s to %s leaves king checked.", piece, to)
	}

	// make move and return captured piece, if any
	capture := game.gameBoard.movePiece(piece, to)

	// Special moves: castling, promoting, etc.
	switch piece.getType() {
	case kingType:
		if rookMove, ok := castlingPositions[to]; ok {
			rook, _ := game.getPlayerPiece(rookMove.rookFrom, player.isWhite)
			game.gameBoard.movePiece(rook, rookMove.rookTo)

			king := player.getKing()

			king.shortCastlingOpt = false
			king.longCastlingOpt = false
		}

	case pawnType:

		// Promotion
		if to.Row == 0 || to.Row == 7 {
			queen := newQueen(to, player)
			game.gameBoard.insertPiece(queen)
			queen.setMoved(true)
			break
		}

		// En passant
		if to.Col != from.Col {
			capturedPos := position{Row: from.Row, Col: to.Col}
			capture, _ = game.gameBoard.getPiece(capturedPos)
			game.gameBoard.clearSquare(capturedPos)
			break
		}

		// Jump
		pawn, _ := castPawn(piece)
		diff := from.Row - to.Row
		pawnJumped := diff == 2 || diff == -2
		if pawnJumped {
			println("pawnjumped")
			player.pawnJumped = pawn
			fmt.Println(player.pawnJumped)
		}

	case rookType:
		king := player.getKing()
		if from.Col == 0 {
			king.shortCastlingOpt = false
		}

		if from.Col == 7 {
			king.longCastlingOpt = false
		}
	}

	// if piece was captured
	if capture != nil {
		opponent.pieces = deletePiece(opponent.pieces, capture)
		player.points += capture.getValue()
	}

	attackedSquares := player.attackedSquares(game.gameBoard)

	// flag as checked or not
	opponent.isChecked = attackedSquares[opponent.king.pos]

	// Check winning / draw conditions
	if !opponent.hasLegalMoves(game) {
		if opponent.isChecked {
			fmt.Println("CHECKMATE!!!", player.name, "WINS")
		} else {
			fmt.Println("Stalemate pal :(")
		}
	} else if len(player.pieces) == 1 && len(opponent.pieces) == 1 {
		fmt.Println("Stalemate pal :(")
	}

	// Update moves history
	move := newMove(piece.clone(), capture, from, to, opponent.isChecked)
	game.moveHistory = append(game.moveHistory, move)

	// Switch turns
	game.WhiteTurn = !game.WhiteTurn
	fmt.Println(game.gameBoard)

	if opponent.pawnJumped != nil {
		opponent.removeJumpedPawn()
	}

	return nil
}

// Returns pointer to player based on color
func (g *game) GetPlayer(pColor bool) *player {
	player := g.pWhite

	if !pColor {
		player = g.pBlack
	}

	return player
}

// Returns pointer to player based on color
func (g *game) GetPlayerCopy(white bool) player {
	player := g.pWhite

	if !white {
		player = g.pBlack
	}

	return *player
}

// Returns pointer to player based on player
func (g *game) GetPlayerOpponent(pColor bool) *player {
	opponent := g.pBlack

	if !pColor {
		opponent = g.pWhite
	}

	return opponent
}

// Returns pointer to player based on player
func (g *game) GetPlayerOpponentCopy(white bool) player {
	opponent := g.pBlack

	if !white {
		opponent = g.pWhite
	}

	return *opponent
}

func (game *game) setupPositionFENString(FENPosition []string) error {

	pWhite := game.pWhite
	pBlack := game.pBlack

	for i, row := range FENPosition {

		colNum := 0
		rowNum := 7 - i

		for _, char := range row {
			// empty squares
			if unicode.IsNumber(char) {
				num := int(char - '0')
				if num < 1 || num > 8 {
					return errors.New("Error reading FEN String: Incorrect format for num in position. Check pos string and try again.")
				}

				colNum += num
				continue
			}

			letter := string(char)
			position := position{Row: rowNum, Col: colNum}
			var piece movable
			switch letter {
			case "p":
				piece = newPawn(position, pBlack)
			case "P":
				piece = newPawn(position, pWhite)
			case "n":
				piece = newKnight(position, pBlack)
			case "N":
				piece = newKnight(position, pWhite)
			case "b":
				piece = newBishop(position, pBlack)
			case "B":
				piece = newBishop(position, pWhite)
			case "r":
				piece = newRook(position, pBlack)
			case "R":
				piece = newRook(position, pWhite)
			case "q":
				piece = newQueen(position, pBlack)
			case "Q":
				piece = newQueen(position, pWhite)
			case "k":
				piece = newKing(position, pBlack)
				king, _ := castKing(piece)
				pBlack.king = king
			case "K":
				piece = newKing(position, pWhite)
				king, _ := castKing(piece)
				pWhite.king = king
			}

			game.gameBoard.insertPiece(piece)
			colNum++
		}

	}

	return nil
}

func (game *game) setupCastlingFENString(FENCastling string) error {

	pWhite := game.pWhite
	pBlack := game.pBlack

	if FENCastling == "-" {
		return nil
	}

	options := []string{"K", "Q", "k", "q"}

	for _, char := range FENCastling {
		letter := string(char)
		if !slices.Contains(options, letter) {
			return fmt.Errorf("Incorrect character for castling options. Is: '%s'. Should be: [ 'K' | 'Q' | 'k' | 'q' ].", letter)
		}
	}

	if !strings.Contains(FENCastling, "K") {
		pWhite.king.shortCastlingOpt = false
	}

	if !strings.Contains(FENCastling, "Q") {
		pWhite.king.longCastlingOpt = false
	}

	if !strings.Contains(FENCastling, "k") {
		pBlack.king.shortCastlingOpt = false
	}

	if !strings.Contains(FENCastling, "q") {
		pBlack.king.longCastlingOpt = false
	}

	return nil
}

func (game *game) setupEnPassantFENString(FENEnPassant string, turn bool) error {

	if FENEnPassant == "-" {
		fmt.Println("No en passant")
		return nil
	}

	capturePosition := pos(FENEnPassant)
	if capturePosition.Col == -1 && capturePosition.Row == -1 {
		return fmt.Errorf("Error reading enpassant target position. Make sure it's the right format and is in bounds of the board.")
	}

	row := capturePosition.getRow()

	if row != 6 && row != 3 {
		return fmt.Errorf("En passant target should be either on rank 6 or 3. Current rank: %b", capturePosition.getRow())
	}

	// If it's white to play, it means the pawnDirection of the pawn is downwards (black pawn).
	// If it's black to play, it means the pawnDirection of the pawn is upwards (white pawn).
	pawnDirection := -1

	if !turn {
		pawnDirection *= -1
	}

	pawnPosition := position{Col: capturePosition.Col, Row: capturePosition.Row + 1*pawnDirection}

	piece, ok := game.gameBoard.getPiece(pawnPosition)
	if !ok {
		return fmt.Errorf("Not pawn found at position %s", pawnPosition)
	}

	pawn, ok := castPawn(piece)
	if !ok {
		return fmt.Errorf("Piece at position %s is not a pawn.", pawnPosition)
	}

	if pawn.isWhite() == turn {
		return fmt.Errorf("Piece at position %s is of same color as player's turn.", pawnPosition)
	}

	pawn.jumped = true

	fmt.Println("en passant for pawn at", pawnPosition.String())

	return nil
}
