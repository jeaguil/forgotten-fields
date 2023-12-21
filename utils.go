package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Vector struct {
	x float64
	y float64
}

const (
	North string = "North"
	South string = "South"
	East  string = "East"
	West  string = "West"
)

type CardinalDirection struct {
	direction string
	angle     float64
}

func IsKeyJustPressed() (bool, ebiten.Key) {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		return true, ebiten.KeyDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		return true, ebiten.KeyUp
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		return true, ebiten.KeyRight
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		return true, ebiten.KeyLeft
	}

	if inpututil.IsKeyJustPressed(ebiten.Key(ebiten.MouseButtonRight)) {
		return true, ebiten.Key(ebiten.MouseButtonRight)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key(ebiten.MouseButtonLeft)) {
		return true, ebiten.Key(ebiten.MouseButtonLeft)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true, ebiten.KeySpace
	}

	return false, -1
}
