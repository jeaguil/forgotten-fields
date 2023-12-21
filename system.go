package main

/*
ECS (Entity Component System) architectural pattern to group components together which carry data
*/

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TransformType ComponentType = "TRANSFORM"
	ImgType       ComponentType = "IMG"
	AnimationType ComponentType = "ANIMATION"
	ScaleType     ComponentType = "SCALE"
)

type TransformComponent struct {
	PosX, PosY float64
}

func (t *TransformComponent) Type() ComponentType { return TransformType }

type ImgComponent struct {
	Image *ebiten.Image
}

func (i *ImgComponent) Type() ComponentType { return ImgType }

type ScaleComponent struct {
	factor int
}

func (s *ScaleComponent) Type() ComponentType { return ScaleType }

type ComponentType string

type ComponentTyper interface{ Type() ComponentType }

type Entity struct {
	components map[ComponentType]ComponentTyper
}

func (e *Entity) HasComponent(cType ComponentType) bool {
	_, ok := e.components[cType]
	return ok
}

func (e *Entity) GetComponent(cType ComponentType) ComponentTyper {
	if c, ok := e.components[cType]; !ok {
		panic(fmt.Sprintf("expected entity to have component of type %s attached", cType))
	} else {
		return c
	}
}

func (e *Entity) AddComponent(c ComponentTyper) {
	if e.HasComponent(c.Type()) {
		panic(fmt.Sprintf("entity already has component of type %s attached", c))
	}
	e.components[c.Type()] = c
}

type Registry struct {
	entites []*Entity
}

func (r *Registry) Query(types ...ComponentType) []*Entity {
	candidates := []*Entity{}

	for _, e := range r.entites {
		matchCount := 0
		for _, c := range types {
			if e.HasComponent(c) {
				matchCount++
			}
		}
		if matchCount == len(types) {
			candidates = append(candidates, e)
		}
	}
	return candidates
}

func (r *Registry) NewEntity() *Entity {
	e := Entity{components: make(map[ComponentType]ComponentTyper)}
	r.entites = append(r.entites, &e)
	return &e
}

type RenderSystem struct {
	Registry *Registry
}

func (r *RenderSystem) Draw(screen *ebiten.Image) {
	for _, e := range r.Registry.Query(TransformType, ImgType) {
		position := e.GetComponent(TransformType).(*TransformComponent)
		img := e.GetComponent(ImgType).(*ImgComponent)
		scale := e.GetComponent(ScaleType).(*ScaleComponent)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.PosX, position.PosY)
		op.GeoM.Scale(float64(scale.factor), float64(scale.factor))
		screen.DrawImage(img.Image, op)
	}
}

type AnimationComponent struct {
	Frames            []*ebiten.Image
	CurrentFrameIndex int
	Count             float64
	AnimationSpeed    float64
}

func (a *AnimationComponent) Type() ComponentType { return AnimationType }

type AnimationSystem struct {
	Registry *Registry
}

func (a *AnimationSystem) Update() error {
	for _, e := range a.Registry.Query(ImgType, AnimationType) {
		a := e.GetComponent(AnimationType).(*AnimationComponent)
		s := e.GetComponent(ImgType).(*ImgComponent)

		a.Count += a.AnimationSpeed
		a.CurrentFrameIndex = int(math.Floor(a.Count))

		if a.CurrentFrameIndex >= len(a.Frames) {
			a.Count = 0
			a.CurrentFrameIndex = 0
		}

		s.Image = a.Frames[a.CurrentFrameIndex]
	}
	return nil
}
