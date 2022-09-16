package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	x int
	y int
}

func MakePoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) getX() int {
	return p.x
}

func (p Point) getY() int {
	return p.y
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x1, y1, x2, y2 := rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)

	p1 := MakePoint(x1, y1)
	p2 := MakePoint(x2, y2)

	fmt.Printf("Point 1: (%v, %v)\n", p1.getX(), p1.getY())
	fmt.Printf("Point 2: (%v, %v)\n", p2.getX(), p2.getY())

	fmt.Printf("\nDistance: %v\n", getDistance(p1, p2))
}

func getDistance(p1, p2 Point) float64 {
	dx := p1.getX() - p2.getX()
	dy := p1.getY() - p2.getY()

	return math.Sqrt(float64(dx*dx + dy*dy))
}
