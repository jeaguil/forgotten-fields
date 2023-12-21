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
	playerIdle *ebiten.Image
	playerWalk *ebiten.Image

	uiFont font.Face

	BackgroundImage *ebiten.Image
)

const (
	playerWidth                  = 16
	playerHeight                 = 64
	playerOX                     = 0
	playerOY                     = 0
	playerFrameCount             = 4
	playerFrameAnimationSequence = 15
	playerFrameWalkSpeed         = 10
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
	uiFont = requiredAssetFont("assets/font.ttf")

	playerIdle = requiredAssetImage("assets/playeridle.png")
	playerWalk = requiredAssetImage("assets/playerwalk.png")

	BackgroundImage = LoadMap()
}
