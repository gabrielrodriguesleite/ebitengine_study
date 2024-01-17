package main

import (
	"log"

	twenty48 "github.com/gabrielrodriguesleite/ebitengine_study/src"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := twenty48.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(twenty48.ScreenWidth, twenty48.ScreenHeight)
	ebiten.SetWindowTitle("My game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
