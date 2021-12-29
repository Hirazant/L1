package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func Distance(point1 Point, point2 Point) float64 {
	dx := point1.x - point2.x
	dy := point1.y - point2.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func main() {
	point1 := NewPoint(7, 7)
	point2 := NewPoint(3, 3)
	fmt.Printf("distance between %v and %v is - %f", point1, point2, Distance(point1, point2))
}
