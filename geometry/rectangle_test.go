package geometry_test

import (
	"math"
	"testing"

	"geomEngine/geometry"
)

func TestRectangleValidation(t *testing.T) {
	_, err := geometry.NewRectangle(5, 4, geometry.Point{X: 0, Y: 0})
	if err != nil {
		t.Fatalf("expected no error for valid circle, got %v", err)
	}

	_, err = geometry.NewRectangle(0, 0, geometry.Point{X: 0, Y: 0})
	if err == nil {
		t.Fatalf("expected error for zero dimension, got nil")
	}
}

func TestRectangleArea(t *testing.T) {
	width := 5.0
	height := 4.0
	c, _ := geometry.NewRectangle(width, height, geometry.Point{X: 0, Y: 0})

	expected := width * height
	actual := c.Area()

	if math.Abs(actual-expected) > epsilon {
		t.Errorf("area mismatch: expected %v, got %v", expected, actual)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	width := 5.0
	height := 4.0
	c, _ := geometry.NewRectangle(width, height, geometry.Point{X: 0, Y: 0})

	expected := 2 * (height + width)
	actual := c.Perimeter()

	if math.Abs(actual-expected) > epsilon {
		t.Errorf("perimeter mismatch: expected %v, got %v", expected, actual)
	}
}

func TestRectangleBoundingBox(t *testing.T) {
	origin := geometry.Point{X: 10, Y: 20}
	width := 5.0
	height := 4.0

	c, _ := geometry.NewRectangle(width, height, origin)
	bb := c.BoundingBox()

	if bb.MinX != origin.X {
		t.Errorf("MinX mismatch")
	}
	if bb.MinY != origin.Y {
		t.Errorf("MinY mismatch")
	}
	if bb.MaxX != origin.X+width {
		t.Errorf("MaxX mismatch")
	}
	if bb.MaxY != origin.Y+height {
		t.Errorf("MaxY mismatch")
	}
}
