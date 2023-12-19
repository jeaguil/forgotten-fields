package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position Vector
	cardinal CardinalDirection
	sprite   *ebiten.Image
}

var (
	isMoving bool = false
)

func NewPlayer() *Player {
	sprite := playerIdle

	// get img size to determine player starting position (will always be center of game screen).
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2
	startingPosition := Vector{
		x: gameScreenWidth/2 - halfW,
		y: gameScreenHeight/2 - halfH,
	}
	startingDirection := CardinalDirection{
		direction: North,
		angle:     -90.0 * math.Pi / 180.0,
	}

	return &Player{
		position: startingPosition,
		cardinal: startingDirection,
		sprite:   sprite,
	}
}

func (p *Player) Update() {
	velocity := 1.5
	isMoving = false
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.position.y += velocity
		p.cardinal.angle = 90.0 * math.Pi / 180.0
		p.cardinal.direction = South
		isMoving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.y -= velocity
		p.cardinal.angle = -90.0 * math.Pi / 180.0
		p.cardinal.direction = North
		isMoving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += velocity
		p.cardinal.angle = 0.0 * math.Pi / 180.0
		p.cardinal.direction = East
		isMoving = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= velocity
		p.cardinal.angle = 180.0 * math.Pi / 180.0
		p.cardinal.direction = West
		isMoving = true
	}
}

func (p *Player) Draw(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	animationSpeed := playerFrameAnimationSequence
	if isMoving {
		p.sprite = playerWalk
		animationSpeed = playerFrameWalkSpeed
	} else {
		p.sprite = playerIdle
	}
	if p.cardinal.direction == West {
		flipSpeed := float64(p.sprite.Bounds().Dx()) / 8
		op.GeoM.Translate(-flipSpeed, 0)
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(flipSpeed, 0)
	}
	op.GeoM.Scale(2.5, 2.5)
	op.GeoM.Translate(p.position.x, p.position.y)
	i := (g.frameCount / animationSpeed) % playerFrameCount
	sx, sy := playerOX+i*playerWidth, 0
	dimensions2 := image.Rect(sx, sy, sx+playerWidth, playerHeight)
	screen.DrawImage(p.sprite.SubImage(dimensions2).(*ebiten.Image), op)
}
