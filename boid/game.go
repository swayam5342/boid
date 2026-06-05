package boid

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	ev "github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/swayam5342/boid/vector"
)

type Boid struct {
	position, vel, target vector.Vec2
}

const (
	moveSpeed                 = 20
	perceptionRadius          = 100
	steerForce                = 1
	alignmentForce            = 0.1
	cohesionForce             = 0.05
	separationForce           = 0.3
	centralizationForce       = 0.3
	centralizationForceRadius = 200
)

func NewBoid(x, y float64, tarcen vector.Vec2) *Boid {
	return &Boid{
		position: vector.Vec2{X: x, Y: y},
		vel:      vector.Vec2{X: rand.Float64()*2 - 1, Y: rand.Float64()*2 - 1},
		target:   tarcen,
	}
}
func (b *Boid) Draw(screen *ebiten.Image) {
	ev.FillCircle(screen, float32(b.position.X), float32(b.position.Y), 10, color.RGBA{255, 148, 148, 0xff}, true)
}

func (b *Boid) Update(boids []*Boid) {
	neighbors := b.getNeighbors(boids)

	alignment := b.alignment(neighbors)
	cohesion := b.cohesion(neighbors)
	separation := b.separation(neighbors)
	centering := b.centralization()

	b.vel = b.vel.Add(alignment).Add(cohesion).Add(separation).Add(centering).Limit(moveSpeed)
	b.position = b.position.Add(b.vel)
}

func (b *Boid) getNeighbors(boids []*Boid) []*Boid {
	var n []*Boid
	for _, other := range boids {
		if b != other && b.position.DistanceTo(other.position) < perceptionRadius {
			n = append(n, other)
		}
	}
	return n
}
func (b *Boid) SetTargetCenter(center vector.Vec2) {
	b.target = center
}
func (b *Boid) alignment(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	if len(boids) == 0 {
		return sum
	}
	for _, other := range boids {
		if b != other {
			sum = sum.Add(other.vel)
		}
	}
	avg := sum.Div(float64(len(boids)))
	return b.steer(avg.Normalize().Mul(moveSpeed)).Mul(alignmentForce)
}

func (b *Boid) cohesion(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	if len(boids) == 0 {
		return sum
	}
	for _, other := range boids {
		if b != other {
			sum = sum.Add(other.position)
		}
	}
	avg := sum.Div(float64(len(boids)))
	return b.steer(avg.Sub(b.position).Normalize().Mul(moveSpeed)).Mul(cohesionForce)
}

func (b *Boid) separation(boids []*Boid) vector.Vec2 {
	var sum vector.Vec2
	var closeN []*Boid
	for _, other := range boids {

		if b != other && b.position.DistanceTo(other.position) < perceptionRadius/2 {
			closeN = append(closeN, other)
		}
	}
	if len(closeN) == 0 {
		return sum
	}
	for _, other := range closeN {
		diff := b.position.Sub(other.position)
		sum = sum.Add(diff.Normalize().Div(diff.Length()))
	}
	avg := sum.Div(float64(len(closeN)))
	return b.steer(avg.Normalize().Mul(moveSpeed)).Mul(separationForce)
}

func (b *Boid) centralization() vector.Vec2 {
	if b.position.DistanceTo(b.target) < centralizationForceRadius {
		return vector.Vec2{}
	}
	d := b.target.Sub(b.position).Normalize().Mul(moveSpeed)
	return b.steer(d).Mul(centralizationForce)
}

func (b *Boid) steer(target vector.Vec2) vector.Vec2 {
	steer := target.Sub(b.vel)
	return steer.Normalize().Mul(steerForce)
}
