package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 420
	ScreenHeight = 600
	boardSize    = 4
)

type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(backGroundColor)
	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() (err error) {
	if err := g.board.Update(
	// input missing
	); err != nil {
		return err
	}
	return nil
}

func NewGame() (*Game, error) {
	g := &Game{
		// input missing
	}

	var err error
	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, err
}
