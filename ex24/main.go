package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Set(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) Get() (float64, float64) {
	return p.x, p.y
}

func main() {
	point1 := &Point{3, 4}
	point2 := &Point{5, 6}

	fmt.Printf("%.2f\n", calculateDistance(point1, point2))
}

func calculateDistance(p1, p2 *Point) float64 {
	x1, y1 := p1.Get()
	x2, y2 := p2.Get()

	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
