package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 420
	ScreenHeight = 600
)

type Game struct {
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() (err error) {
	return
}

func NewGame() (*Game, error) {
	g := &Game{}
	var err error

	return g, err
}
