package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Forgotten-Fields")
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizable(true)
	g, err := StartGame()
	if err != nil {
		log.Fatal(err)
	}

	if err = ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
