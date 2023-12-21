package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gameScreenWidth  = 960
	gameScreenHeight = 640

	gameTitle = "Forgotten Fields"
)

func init() {
	LoadGameAssets()
}

func main() {
	ebiten.SetWindowTitle(gameTitle)
	ebiten.SetWindowSize(gameScreenWidth, gameScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)

	g, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err = ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
