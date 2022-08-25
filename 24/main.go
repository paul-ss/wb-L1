package main

import (
	"fmt"
	"math"
)

func main() {
	p1, p2 := NewPoint(0.0, 0.0), NewPoint(-3.0, 4.0)
	fmt.Println(Distance(p1, p2))
}

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
