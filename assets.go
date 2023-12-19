package main

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed assets/*
var assets embed.FS

var (
	// main player sprite
	playerCar  *ebiten.Image
	playerIdle *ebiten.Image

	uiFont font.Face
)

const (
	playerWidth          = 16
	playerHeight         = 64
	playerOX             = 0
	playerOY             = 0
	playerFrameCount     = 4
	playerAnimationSpeed = 10
)

func requiredAssetImage(filepath string) *ebiten.Image {
	f, err := assets.Open(filepath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func requiredAssetFont(filepath string) font.Face {
	f, err := assets.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}

func LoadGameAssets() {
	playerCar = requiredAssetImage("assets/kenney_robotpack/robot_3Dred.png")
	playerIdle = requiredAssetImage("assets/playeridle.png")
	uiFont = requiredAssetFont("assets/font.ttf")
}
