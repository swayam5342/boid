package boid

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/swayam5342/boid/vector"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
	boidCount    = 100
)

type Game struct {
	boids []*Boid
}

func NewGame() (*Game, error) {
	g := &Game{
		boids: make([]*Boid, boidCount),
	}

	for i := range g.boids {
		g.boids[i] = NewBoid(
			rand.Float64()*ScreenWidth,
			rand.Float64()*ScreenHeight,
			vector.Vec2{X: ScreenWidth / 2, Y: ScreenHeight / 2},
		)
	}
	return g, nil
}

func (g *Game) Update() error {
	for _, b := range g.boids {
		b.Update(g.boids)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 245, 228, 0xff})
	for _, b := range g.boids {
		b.Draw(screen)
	}
	fps := fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, fps)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
