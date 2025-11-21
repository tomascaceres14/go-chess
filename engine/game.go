package gochess

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
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
	id                           string
	gameBoard                    *board
	pWhite, pBlack               *player
	WhiteTurn                    bool
	halfmoveClock, fullmoveCount int
	moveHistory                  []move
	status                       gameStatus
	castleDir                    int
}

func newGame(whiteName, blackName string) (*game, error) {

	gameBoard := [8][8]Movable{}

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
		fullmoveCount: 1,
		halfmoveClock: 0,
		castleDir:     -1,
	}

	return game, nil
}

// Generates a new board with classic chess configuration
func newGameClassic(whiteName, blackName string) (*game, error) {

	game, err := newGame(whiteName, blackName)
	if err != nil {
		return nil, err
	}

	pWhite := game.pWhite
	pBlack := game.pBlack

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

	return game, nil
}

func newGameFENString(FENString string, whiteName, blackName string) (*game, error) {

	game, err := newGame(whiteName, blackName)
	if err != nil {
		return nil, err
	}

	FENSplit := strings.Split(FENString, " ")
	if len(FENSplit) != 6 {
		return nil, errors.New("Error reading FEN String: Incorrect format. FEN string should contain 6 tramos")
	}

	FENPosition := strings.Split(FENSplit[0], "/")
	if len(FENPosition) != 8 {
		return nil, errors.New("Error reading FEN String: Incorrect format for position. Should contain 8 tramos")
	}

	if err := game.setFENStringPos(FENPosition); err != nil {
		return nil, err
	}

	FENTurn := FENSplit[1]
	if len(FENTurn) != 1 || (FENTurn != "w" && FENTurn != "b") {
		return nil, errors.New("Incorrect format for turn. Should be 1 character, either 'b' or 'w'.")
	}

	game.WhiteTurn = FENTurn == "w"

	FENCastling := FENSplit[2]
	if err := game.setFENStringCastling(FENCastling); err != nil {
		return nil, err
	}

	FENEnPassant := FENSplit[3]
	if err := game.setFENStringEnPassant(FENEnPassant); err != nil {
		return nil, err
	}

	FENHalfmoveClock := FENSplit[4]
	halfmoveClock, err := strconv.Atoi(FENHalfmoveClock)
	if err != nil {
		return nil, fmt.Errorf("Halfmove clock should be a number. Got: %s", FENHalfmoveClock)
	}

	if err := game.setHalfmoveClock(halfmoveClock); err != nil {
		return nil, err
	}

	FENFullmoveCount := FENSplit[5]
	fullmoveCount, err := strconv.Atoi(FENFullmoveCount)
	if err != nil {
		return nil, fmt.Errorf("Fullmove count should be a number. Got: %s", FENHalfmoveClock)
	}

	if err := game.setFullmoveCount(fullmoveCount); err != nil {
		return nil, err
	}

	return game, nil
}

func (g *game) validateMove(move *move) error {
	color := move.color
	from := move.from
	to := move.to

	// Check turn to play
	if g.WhiteTurn != color {
		return fmt.Errorf("Not your turn, %s.", colorToString(color))
	}

	// Check if positions are in bounds
	if !from.inBounds() || !to.inBounds() {
		return errors.New("Position out of bounds.")
	}

	// Obtains piece to move
	piece, err := g.getPlayerPiece(from, color)
	if err != nil {
		return err
	}

	move.piece = piece

	// Check if piece can move to desired position or if is trying to move in-place
	legalMoves := piece.legalMoves(g.gameBoard)
	if !legalMoves[to] || piece.getPosition() == to {
		return fmt.Errorf("%s can't move from %s to %s.", piece.String(), from, to)
	}

	// Check if the move leaves king vulnerable
	if !isMoveSafeToKing(piece, to, g.gameBoard) {
		return fmt.Errorf("%s to %s leaves king checked.", piece, to)
	}

	return nil
}

// Moves piece in position `from` to position `to` if player is owner of piece
func (g *game) makeMove(from, to position, pColor bool) error {

	move := move{
		from:      from,
		to:        to,
		color:     pColor,
		castleDir: -1,
	}

	if err := g.validateMove(&move); err != nil {
		return err
	}

	piece := move.getPiece()

	// make move and return captured piece, if any
	capture := piece.move(to, g)

	// Obtain player and opponent
	player := g.GetPlayer(pColor)
	opponent := g.GetPlayerOpponent(player.isWhite)

	// if piece was captured
	if capture != nil {
		opponent.deletePiece(capture)
		player.incrementPoints(capture.getValue())
		move.capture = capture
	}

	isCheck := g.isKingInCheck(!pColor)

	// flag as checked or not
	opponent.isChecked = isCheck
	move.isCheck = isCheck

	// Update moves history
	g.moveHistory = append(g.moveHistory, move)

	// Check winning / draw conditions
	if !opponent.hasLegalMoves(g) {
		if opponent.isChecked {
			fmt.Println("CHECKMATE!!!", player.name, "WINS")
		} else {
			fmt.Println("Stalemate pal :(")
		}
	} else if len(player.pieces) == 1 && len(opponent.pieces) == 1 {
		fmt.Println("Stalemate pal :(")
	}

	g.switchTurns()

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

// Obtains piece at given position if player is owner of piece
func (g *game) getPlayerPiece(pos position, pColor bool) (Movable, error) {

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

func (g *game) isKingInCheck(color bool) bool {
	player := g.GetPlayer(color)
	kingPos := player.getKing().getPosition()
	return g.gameBoard.isKingInCheck(kingPos, color)
}

func (g *game) switchTurns() {

	g.WhiteTurn = !g.WhiteTurn
	g.castleDir = -1
	g.fullmoveCount++
}

func (g *game) GetFENString() string {

	FENString := getFENPosition(g)

	// Define turn of player
	turn := " w "
	if !g.WhiteTurn {
		turn = " b "
	}

	FENString += turn
	FENString += getFENCastling(g)
	FENString += getFENEnPassant(g)
	FENString += strconv.Itoa(g.halfmoveClock) + " "
	FENString += strconv.Itoa(g.fullmoveCount)

	return FENString
}

func (g *game) setFENStringPos(FENPosition []string) error {

	pWhite := g.pWhite
	pBlack := g.pBlack

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
			position := position{row: rowNum, col: colNum}
			var piece Movable
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

			g.gameBoard.insertPiece(piece)
			colNum++
		}

	}

	return nil
}

func (g *game) setFENStringCastling(FENCastling string) error {

	pWhite := g.pWhite
	pBlack := g.pBlack

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

func (g *game) setFENStringEnPassant(FENEnPassant string) error {

	if FENEnPassant == "-" {
		return nil
	}

	turn := g.WhiteTurn

	capturePosition := pos(FENEnPassant)
	if capturePosition.col == -1 && capturePosition.row == -1 {
		return fmt.Errorf("Error reading enpassant target position. Make sure it's the right format and is in bounds of the board.")
	}

	row := capturePosition.getRank()
	if row != 6 && row != 3 {
		return fmt.Errorf("En passant target should be either on rank 6 or 3. Current rank: %d", capturePosition.getRank())
	}

	// If it's white to play, it means the pawnDirection of the pawn is downwards (black pawn).
	// If it's black to play, it means the pawnDirection of the pawn is upwards (white pawn).
	pawnDirection := -1
	if !turn {
		pawnDirection *= -1
	}

	pawnPosition := position{col: capturePosition.col, row: capturePosition.row + 1*pawnDirection}

	piece, ok := g.gameBoard.getPiece(pawnPosition)
	if !ok {
		return fmt.Errorf("No piece found at position %s", pawnPosition)
	}

	pawn, ok := castPawn(piece)
	if !ok {
		return fmt.Errorf("Piece at position %s is not a pawn.", pawnPosition)
	}

	if pawn.isWhite() == turn {
		return fmt.Errorf("Pawn at %s is of same color as player's turn.", pawnPosition)
	}

	g.gameBoard.enPassantTarget = &capturePosition

	return nil
}

func (g *game) setHalfmoveClock(halfmoveClock int) error {

	if 0 > halfmoveClock || halfmoveClock > 50 {
		return fmt.Errorf("Halfmove clock should be a number between 0 and 50. Got: %d", halfmoveClock)
	}

	g.halfmoveClock = halfmoveClock

	return nil
}

func (g *game) setFullmoveCount(fullmoveCount int) error {

	if 1 > fullmoveCount {
		return fmt.Errorf("Halfmove clock should be a number greater than 1. Got: %d", fullmoveCount)
	}

	g.fullmoveCount = fullmoveCount

	return nil
}
