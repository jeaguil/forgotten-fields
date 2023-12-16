package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	x float64
	y float64
}

type Player struct {
	position Vector
	cardinal CardinalDirection
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := PlayerCar

	// get img size to determine player starting position (will always be center of game screen).
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2
	startingPosition := Vector{
		x: gameScreenWidth/2 - halfW,
		y: gameScreenHeight/2 - halfH,
	}
	startingDirection := CardinalDirection{
		direction: 0,
		angle:     -90.0 * math.Pi / 180.0,
	}

	return &Player{
		position: startingPosition,
		cardinal: startingDirection,
		sprite:   PlayerCar,
	}
}

type CardinalDirection struct {
	direction int
	angle     float64
}

func (p *Player) Update() {
	velocity := 2.5
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.position.y += velocity
		p.cardinal.angle = 90.0 * math.Pi / 180.0
		p.cardinal.direction = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.y -= velocity
		p.cardinal.angle = -90.0 * math.Pi / 180.0
		p.cardinal.direction = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += velocity
		p.cardinal.angle = 0.0 * math.Pi / 180.0
		p.cardinal.direction = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= velocity
		p.cardinal.angle = 180.0 * math.Pi / 180.0
		p.cardinal.direction = 3
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2
	drawOptions.GeoM.Translate(-halfW, -halfH)
	drawOptions.GeoM.Rotate(p.cardinal.angle)
	drawOptions.GeoM.Translate(halfW, halfH)
	drawOptions.GeoM.Translate(p.position.x, p.position.y)
	screen.DrawImage(p.sprite, drawOptions)
}
