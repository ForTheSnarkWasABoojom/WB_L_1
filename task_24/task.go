package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	pointA := NewPoint(10.0, -22.0)
	pointB := NewPoint(-6.0, 16.0)

	distance := pointA.Distance(pointB)

	fmt.Printf("Distance between point A and point B is %.2f\n", distance)
}
