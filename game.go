package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	playerStartingPosition Vector
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	drawOptions := &ebiten.DrawImageOptions{}

	// get img size to determine player starting position (will always be center of game screen).
	x := PlayerCar.Bounds().Dx()
	y := PlayerCar.Bounds().Dy()
	drawOptions.GeoM.Translate(g.playerStartingPosition.x-float64(x)/2, g.playerStartingPosition.y-float64(y)/2)

	screen.DrawImage(PlayerCar, drawOptions)
}

// Layout is called when the Game's layout changes.
func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

// NewGame returns a new Forgotten-Fields Game.
func NewGame() (*Game, error) {
	g := &Game{
		playerStartingPosition: Vector{x: gameScreenWidth / 2, y: gameScreenHeight / 2},
	}
	return g, nil
}
