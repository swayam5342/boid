package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/swayam5342/boid/boid"
)

func main() {
	game, err := boid.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(boid.ScreenWidth, boid.ScreenHeight)
	ebiten.SetFullscreen(false)
	ebiten.SetWindowTitle("Boids")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
