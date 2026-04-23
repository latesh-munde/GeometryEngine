package geometry

import "math"

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func IsEqualPoints(p1, p2 Point) bool {
	dx := math.Abs(p1.X - p2.X)
	dy := math.Abs(p1.Y - p2.Y)
	if (dx < 10e-6) && (dy < 10e-6) {
		return true
	}
	return false
}
