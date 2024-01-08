package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (

	// Define game movement boundaries
	minX = 0.0
	maxX = float64(gameScreenWidth - 16.0) // Player image width is 16

	minY = 0.0
	maxY = float64(gameScreenHeight - 16.0) // Player image height is 16
)

type Game struct {
	player       *Player
	frameCount   int
	RenderSystem *RenderSystem
}

func (g *Game) Update() error {
	g.player.Update()
	g.frameCount++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(BackgroundImage, nil)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	text.Draw(
		screen,
		gameTitle,
		uiFont,
		gameScreenWidth/2-200,
		50,
		color.White)

	text.Draw(
		screen,
		fmt.Sprintf("%v", g.player.cardinal.direction),
		uiFont,
		50,
		gameScreenHeight-50,
		color.White,
	)
	g.player.Draw(screen, g)
	g.RenderSystem.Draw(screen)
}

// Layout is called when the Game's layout changes.
func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

// NewGame returns a new Forgotten-Fields Game.
func NewGame() (*Game, error) {
	registry := Registry{}
	renderSystem := RenderSystem{Registry: &registry}
	g := &Game{
		RenderSystem: &renderSystem,
		player:       NewPlayer(),
	}
	return g, nil
}
