package geometry

import (
	"geomEngine/errors"
	"math"
)

type Circle struct {
	Center Point
	Radius float64
}

func NewCircle(radius float64, center Point) (Circle, error) {
	if radius <= 0 {
		return Circle{}, errors.ErrInvalidDimension
	}
	return Circle{Radius: radius, Center: center}, nil
}

func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}

func (c Circle) Perimeter() float64 {
	return (2 * math.Pi * c.Radius)
}

func (c Circle) BoundingBox() BoundingBox {
	return NewBoundingBox(
		c.Center.X-c.Radius,
		c.Center.Y-c.Radius,
		c.Center.X+c.Radius,
		c.Center.Y+c.Radius,
	)
	// MinX = cx - r
	// MinY = cy - r
	// MaxX = cx + r
	// MaxY = cy + r
}

func (c Circle) Intersects(other Shape) (bool, error) {
	return false, nil
}
