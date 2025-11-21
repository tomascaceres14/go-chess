package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	gochess "github.com/tomascaceres14/go-chess/engine"
)

const (
	tileSize  = 130
	boardSize = 8
)

type Sprite struct {
	Img        *ebiten.Image
	posX, posY float64
}

// return sprite of img centered in indicated x and y positions
func CenteredPieceSprite(img *ebiten.Image, x, y float64) *Sprite {
	imgLength := img.Bounds().Dx()
	imgWidth := img.Bounds().Dy()
	return &Sprite{
		Img:  img,
		posX: x + float64((tileSize-imgLength)/2),
		posY: y + float64((tileSize-imgWidth)/2),
	}
}

type Piece struct {
	*Sprite
	IsWhite bool
	piece   gochess.Movable
}

func NewPiece(img *ebiten.Image, imgX, imgY float64, isWhite bool) *Piece {
	return &Piece{
		Sprite:  CenteredPieceSprite(img, imgX, imgY),
		IsWhite: isWhite,
	}
}

// cuts img and extracts piece sub image
func PieceSubImg(img *ebiten.Image, x0, y0, x1, y1 int) *ebiten.Image {
	return img.SubImage(
		image.Rect(x0, y0, x1, y1),
	).(*ebiten.Image)
}

func NewPawn(isWhite bool, img *ebiten.Image, x, y float64) *Piece {

	X0 := 278
	Y0 := 136
	X1 := 356
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	return NewPiece(subImg, x, y, isWhite)
}

func NewBishop(isWhite bool, img *ebiten.Image, x, y float64) *Piece {

	X0 := 127
	Y0 := 2
	X1 := 226
	Y1 := 117

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	return NewPiece(subImg, x, y, isWhite)

}

func NewKnight(isWhite bool, img *ebiten.Image, x, y float64) *Piece {

	X0 := 129
	Y0 := 140
	X1 := 218
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	return NewPiece(subImg, x, y, isWhite)

}

func NewRook(isWhite bool, img *ebiten.Image, x, y float64) *Piece {

	x0 := 0
	y0 := 0
	x1 := 82
	y1 := 117

	subImg := PieceSubImg(img, x0, y0, x1, y1)

	return NewPiece(subImg, x, y, isWhite)
}

func NewQueen(isWhite bool, img *ebiten.Image, x, y float64) *Piece {
	X0 := 3
	Y0 := 136
	X1 := 79
	Y1 := 256

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	return NewPiece(subImg, x, y, isWhite)
}

func NewKing(isWhite bool, img *ebiten.Image, x, y float64) *Piece {
	X0 := 279
	Y0 := 0
	X1 := 356
	Y1 := 117

	subImg := PieceSubImg(img, X0, Y0, X1, Y1)

	return NewPiece(subImg, x, y, isWhite)

}

type Player struct {
	Pieces []*Piece
}

type Game struct {
	WPlayer            *Player
	BPlayer            *Player
	paddingX, paddingY float32
	board              Board
}

func (g *Game) Update() error {

	for _, v := range g.WPlayer.Pieces {
		// reacts to key presses
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			v.posX += 2
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			v.posX -= 2
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			v.posY -= 2
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			v.posY += 2
		}
	}

	for _, v := range g.BPlayer.Pieces {
		// reacts to key presses
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			v.posX += 2
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			v.posX -= 2
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			v.posY -= 2
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			v.posY += 2
		}
	}

	return nil
}

type Board struct {
	grid [8][8]*Piece
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	// draw board
	for i := range boardSize {
		for j := range boardSize {
			if (i+j)%2 == 0 {
				vector.FillRect(
					screen,
					g.paddingX+float32(i*tileSize),
					g.paddingY+float32(j*tileSize),
					tileSize,
					tileSize,
					color.RGBA{240, 217, 181, 255},
					false,
				)
			} else {
				vector.FillRect(
					screen,
					g.paddingX+float32(i*tileSize),
					g.paddingY+float32(j*tileSize),
					tileSize,
					tileSize,
					color.RGBA{181, 136, 99, 255},
					false,
				)
			}
		}
	}

	opts := &ebiten.DrawImageOptions{}

	for _, v := range g.WPlayer.Pieces {
		opts.GeoM.Translate(v.posX, v.posY)

		screen.DrawImage(
			v.Img,
			opts,
		)
		opts.GeoM.Reset()
	}

	for _, v := range g.BPlayer.Pieces {
		opts.GeoM.Translate(v.posX, v.posY)

		screen.DrawImage(
			v.Img,
			opts,
		)
		opts.GeoM.Reset()
	}

	ebitenutil.DebugPrint(screen, "Hello, World!")
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	ebiten.SetWindowTitle("Mates")

	g := &Game{}

	winX, winY := ebiten.WindowSize()
	g.paddingX = float32(winX)/2 - (tileSize*boardSize)/2
	g.paddingY = float32(winY)/2 - (tileSize*boardSize)/2

	whitePieces, _, err := ebitenutil.NewImageFromFile(`assets/images/whitepieces.png`)
	if err != nil {
		log.Fatal(err)
	}

	blackPieces, _, err := ebitenutil.NewImageFromFile(`assets/images/blackpieces.png`)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: pls make func to generate automatically positions based on algebraic notation.
	// This burns my eyes.
	wPlayer := &Player{
		Pieces: []*Piece{
			NewRook(true, whitePieces, float64(g.paddingX), float64(g.paddingY)+tileSize*7),
			NewKnight(true, whitePieces, float64(g.paddingX)+tileSize, float64(g.paddingY)+tileSize*7),
			NewBishop(true, whitePieces, float64(g.paddingX)+tileSize*2, float64(g.paddingY)+tileSize*7),
			NewQueen(true, whitePieces, float64(g.paddingX)+tileSize*3, float64(g.paddingY)+tileSize*7),
			NewKing(true, whitePieces, float64(g.paddingX)+tileSize*4, float64(g.paddingY)+tileSize*7),
			NewBishop(true, whitePieces, float64(g.paddingX)+tileSize*5, float64(g.paddingY)+tileSize*7),
			NewKnight(true, whitePieces, float64(g.paddingX)+tileSize*6, float64(g.paddingY)+tileSize*7),
			NewRook(true, whitePieces, float64(g.paddingX)+tileSize*7, float64(g.paddingY)+tileSize*7),
			NewPawn(true, whitePieces, float64(g.paddingX), float64(g.paddingY)+tileSize*6),
			NewPawn(true, whitePieces, float64(g.paddingX)+tileSize, float64(g.paddingY)+tileSize*6),
			NewPawn(true, whitePieces, float64(g.paddingX)+tileSize*2, float64(g.paddingY)+tileSize*6),
			NewPawn(true, whitePieces, float64(g.paddingX)+tileSize*3, float64(g.paddingY)+tileSize*6),
			NewPawn(true, whitePieces, float64(g.paddingX)+tileSize*4, float64(g.paddingY)+tileSize*6),
			NewPawn(true, whitePieces, float64(g.paddingX)+tileSize*5, float64(g.paddingY)+tileSize*6),
			NewPawn(true, whitePieces, float64(g.paddingX)+tileSize*6, float64(g.paddingY)+tileSize*6),
			NewPawn(true, whitePieces, float64(g.paddingX)+tileSize*7, float64(g.paddingY)+tileSize*6),
		},
	}

	bPlayer := &Player{
		Pieces: []*Piece{
			NewRook(false, blackPieces, float64(g.paddingX), float64(g.paddingY)),
			NewKnight(false, blackPieces, float64(g.paddingX)+tileSize, float64(g.paddingY)),
			NewBishop(false, blackPieces, float64(g.paddingX)+tileSize*2, float64(g.paddingY)),
			NewQueen(false, blackPieces, float64(g.paddingX)+tileSize*3, float64(g.paddingY)),
			NewKing(false, blackPieces, float64(g.paddingX)+tileSize*4, float64(g.paddingY)),
			NewBishop(false, blackPieces, float64(g.paddingX)+tileSize*5, float64(g.paddingY)),
			NewKnight(false, blackPieces, float64(g.paddingX)+tileSize*6, float64(g.paddingY)),
			NewRook(false, blackPieces, float64(g.paddingX)+tileSize*7, float64(g.paddingY)),
			NewPawn(false, blackPieces, float64(g.paddingX), float64(g.paddingY)+tileSize),
			NewPawn(false, blackPieces, float64(g.paddingX)+tileSize, float64(g.paddingY)+tileSize),
			NewPawn(false, blackPieces, float64(g.paddingX)+tileSize*2, float64(g.paddingY)+tileSize),
			NewPawn(false, blackPieces, float64(g.paddingX)+tileSize*3, float64(g.paddingY)+tileSize),
			NewPawn(false, blackPieces, float64(g.paddingX)+tileSize*4, float64(g.paddingY)+tileSize),
			NewPawn(false, blackPieces, float64(g.paddingX)+tileSize*5, float64(g.paddingY)+tileSize),
			NewPawn(false, blackPieces, float64(g.paddingX)+tileSize*6, float64(g.paddingY)+tileSize),
			NewPawn(false, blackPieces, float64(g.paddingX)+tileSize*7, float64(g.paddingY)+tileSize),
		},
	}

	g.WPlayer = wPlayer
	g.BPlayer = bPlayer

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
