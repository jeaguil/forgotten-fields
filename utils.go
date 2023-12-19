package main

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
