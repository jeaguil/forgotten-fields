package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	g.player.Update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen, g)
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
