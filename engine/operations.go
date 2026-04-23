package engine

// Operation represents a geometry operation that can be executed on a Shape.
// This is intentionally small and enum-like.
type Operation int

const (
	// OpArea computes the area of a shape
	OpArea Operation = iota
	// OpPerimeter computes the perimeter/length of a shape
	OpPerimeter
	// OpBoundingBox computes the axis-aligned bounding box of a shape
	OpBoundingBox
)
