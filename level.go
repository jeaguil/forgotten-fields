package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

const (
	mapPath string = "./assets/tiles/forgotten-fields.tmx"
)

func LoadMap() *ebiten.Image {
	levelMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		log.Fatalf("error parsing map: %s", err.Error())
		panic(err)
	}
	renderer, err := render.NewRenderer(levelMap)
	if err != nil {
		log.Fatalf("map unsupported for rendering: %s", err.Error())
		panic(err)
	}
	err = renderer.RenderVisibleLayers()
	if err != nil {
		log.Fatalf("layer unsupported for rendering: %s", err.Error())
		panic(err)
	}
	return ebiten.NewImageFromImage(renderer.Result)
}
