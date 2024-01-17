package main

import (
	"log"

	game "github.com/gabrielrodriguesleite/ebitengine_study/src"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	// ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("My game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
