package geometry

import "geomEngine/errors"

type Rectangle struct {
	Origin        Point // bottom - left corner
	Width, Height float64
}

func NewRectangle(width, height float64, origin Point) (Rectangle, error) {
	if width <= 0 || height <= 0 {
		return Rectangle{}, errors.ErrInvalidDimension
	}
	return Rectangle{Width: width, Height: height, Origin: origin}, nil
}

func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

func (r Rectangle) Perimeter() float64 {
	return (2 * (r.Width + r.Height))
}

func (r Rectangle) BoundingBox() BoundingBox {
	return NewBoundingBox(
		r.Origin.X,
		r.Origin.Y,
		r.Origin.X+r.Width,
		r.Origin.Y+r.Height,
	)
}

func (r Rectangle) Intersects(other Shape) (bool, error) {
	return false, nil
}
