package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(
		screen,
		fmt.Sprintf("%v", g.player.cardinal.direction),
		UI_FONT,
		gameScreenWidth/2,
		50,
		color.White)
	g.player.Draw(screen)
}

// Layout is called when the Game's layout changes.
func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

// NewGame returns a new Forgotten-Fields Game.
func NewGame() (*Game, error) {
	g := &Game{}
	g.player = NewPlayer()
	return g, nil
}
