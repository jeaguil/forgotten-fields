package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Camera struct {
	color   color.Color
	radius  float64
	scale   int
	options ebiten.DrawImageOptions
	X       float64
	Y       float64
}

func (c *Camera) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(c.X), float32(c.Y), float32(c.radius), float32(c.radius), c.color, false)
}

func (c *Camera) Update(p *Player) {
	// Draw circle around player's position
	circleX, circleY := p.position.x, p.position.y
	opCircle := &ebiten.DrawImageOptions{}
	opCircle.GeoM.Translate(-c.radius, -c.radius)
	opCircle.GeoM.Scale(float64(c.scale), float64(c.scale))
	translateOpX, translateOpY := float64(gameScreenWidth/(2*c.scale)), float64(gameScreenHeight/(2*c.scale))
	opCircle.GeoM.Translate(translateOpX, translateOpY)
	opCircle.GeoM.Translate(circleX, circleY)
	c.options = *opCircle
	c.X = circleX
	c.Y = circleY
}
