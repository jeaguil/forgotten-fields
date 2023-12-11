package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gameScreenWidth  = 500
	gameScreenHeight = 500
)

func init() {}

func main() {
	ebiten.SetWindowTitle("Forgotten-Fields")
	ebiten.SetWindowSize(gameScreenWidth, gameScreenHeight)
	ebiten.SetWindowResizable(false)

	g, err := NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err = ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
