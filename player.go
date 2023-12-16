package main

import "github.com/hajimehoshi/ebiten/v2"

type Vector struct {
	x float64
	y float64
}

type Player struct {
	position Vector
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := PlayerCar

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	startingPosition := Vector{
		x: gameScreenWidth/2 - halfW,
		y: gameScreenHeight/2 - halfH,
	}

	return &Player{
		position: startingPosition,
		sprite:   PlayerCar,
	}
}

func (p *Player) Update(g *Game) {
	speed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.position.y += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.position.y -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.position.x += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.position.x -= speed
	}
}

func (p *Player) Draw(screen *ebiten.Image, g *Game) {
	drawOptions := &ebiten.DrawImageOptions{}

	// get img size to determine player starting position (will always be center of game screen).
	x := PlayerCar.Bounds().Dx()
	y := PlayerCar.Bounds().Dy()
	drawOptions.GeoM.Translate(g.player.position.x-float64(x)/2, g.player.position.y-float64(y)/2)

	screen.DrawImage(PlayerCar, drawOptions)
}
