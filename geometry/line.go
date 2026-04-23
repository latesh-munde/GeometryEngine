package geometry

import (
	"geomEngine/errors"
	"math"
)

type Line struct {
	StartPt, EndPt Point
}

func NewLine(start, end Point) (Line, error) {
	if IsEqualPoints(start, end) {
		return Line{}, errors.ErrLine
	}
	return Line{StartPt: start, EndPt: end}, nil
}

func (l Line) Area() float64 {
	return 0.0
}

func (l Line) Perimeter() float64 {
	return Length(l.StartPt, l.EndPt)
}

func (l Line) BoundingBox() BoundingBox {
	return NewBoundingBox(
		math.Min(l.StartPt.X, l.EndPt.X),
		math.Min(l.StartPt.Y, l.EndPt.Y),
		math.Max(l.StartPt.X, l.EndPt.X),
		math.Max(l.StartPt.Y, l.EndPt.Y),
	)
}

func (l Line) Intersects(other Shape) (bool, error) {
	return false, nil
}

func Length(start, end Point) float64 {
	xSq := math.Pow((end.X - start.X), 2)
	ySq := math.Pow((end.Y - start.Y), 2)

	return math.Sqrt(xSq + ySq)
}
