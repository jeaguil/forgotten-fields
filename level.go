package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

const (
	level0FilePath string = "./assets/tiles/forgotten-fields.tmx"
)

type Level struct {
	pathLayerImg *image.NRGBA
	mapImg       *ebiten.Image
}

func (l *Level) loadMap() {
	levelMap, err := tiled.LoadFile(level0FilePath)
	if err != nil {
		log.Fatalf("error parsing map: %s", err.Error())
		panic(err)
	}
	renderer, err := render.NewRenderer(levelMap)
	if err != nil {
		log.Fatalf("map unsupported for rendering: %s", err.Error())
		panic(err)
	}
	err = renderer.RenderGroup(0)
	if err != nil {
		log.Fatalf("error rendering path template layers: %s", err.Error())
		panic(err)
	}
	l.pathLayerImg = renderer.Result
	renderer.Clear()

	err = renderer.RenderVisibleGroups()
	if err != nil {
		log.Fatalf("group unsupported for rendering: %s", err.Error())
		panic(err)
	}
	l.mapImg = ebiten.NewImageFromImage(renderer.Result)
}
