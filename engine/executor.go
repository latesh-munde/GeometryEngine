package engine

import (
	"geomEngine/geometry"
)

func Execute(shape geometry.Shape, op Operation) (any, error) {

	switch op {
	case OpArea:
		return shape.Area(), nil
	case OpPerimeter:
		return shape.Perimeter(), nil
	case OpBoundingBox:
		return shape.BoundingBox(), nil
	default:
		return nil, ErrUnknownOperation
	}
}
