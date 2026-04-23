package geometry

import (
	"geomEngine/errors"
	"math"
)

type Polygon struct {
	Points []Point
}

func NewPolygon(points []Point) (Polygon, error) {
	if len(points) < 3 {
		return Polygon{}, errors.ErrInvalidPolygon
	}
	return Polygon{Points: points}, nil
}

func (p Polygon) Area() float64 {
	length := len(p.Points)

	sum := 0.0

	for i := 0; i < (length - 1); i++ {
		sum += ((p.Points[i].X) * (p.Points[i+1].Y))
		sum -= ((p.Points[i+1].X) * (p.Points[i].Y))
	}
	sum += (p.Points[length-1].X) * (p.Points[0].Y)
	sum -= (p.Points[0].X) * (p.Points[length-1].Y)

	return 0.5 * math.Abs(sum)
}

func (p Polygon) Perimeter() float64 {
	length := len(p.Points)
	perimeter := 0.0

	for i := 0; i < (length - 1); i++ {
		perimeter += Length(p.Points[i], p.Points[i+1])
	}
	perimeter += Length(p.Points[length-1], p.Points[0])

	return perimeter
}

func (p Polygon) BoundingBox() BoundingBox {
	minX := p.Points[0].X
	minY := p.Points[0].Y
	maxX := p.Points[0].X
	maxY := p.Points[0].Y

	for _, val := range p.Points {
		if val.X < minX {
			minX = val.X
		}
		if val.Y < minY {
			minY = val.Y
		}

		if val.X > maxX {
			maxX = val.X
		}
		if val.Y > maxY {
			maxY = val.Y
		}
	}

	return NewBoundingBox(minX, minY, maxX, maxY)
}

func (p Polygon) Intersects(other Shape) (bool, error) {
	return false, nil
}
