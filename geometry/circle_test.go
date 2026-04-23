package geometry_test

import (
	"math"
	"testing"

	"geomEngine/geometry"
)

const epsilon = 0.0001

func TestCircleValidation(t *testing.T) {
	_, err := geometry.NewCircle(5, geometry.Point{X: 0, Y: 0})
	if err != nil {
		t.Fatalf("expected no error for valid circle, got %v", err)
	}

	_, err = geometry.NewCircle(0, geometry.Point{X: 0, Y: 0})
	if err == nil {
		t.Fatalf("expected error for zero radius, got nil")
	}
}

func TestCircleArea(t *testing.T) {
	radius := 2.0
	c, _ := geometry.NewCircle(radius, geometry.Point{X: 0, Y: 0})

	expected := math.Pi * radius * radius
	actual := c.Area()

	if math.Abs(actual-expected) > epsilon {
		t.Errorf("area mismatch: expected %v, got %v", expected, actual)
	}
}

func TestCirclePerimeter(t *testing.T) {
	radius := 3.0
	c, _ := geometry.NewCircle(radius, geometry.Point{X: 0, Y: 0})

	expected := 2 * math.Pi * radius
	actual := c.Perimeter()

	if math.Abs(actual-expected) > epsilon {
		t.Errorf("perimeter mismatch: expected %v, got %v", expected, actual)
	}
}

func TestCircleBoundingBox(t *testing.T) {
	center := geometry.Point{X: 10, Y: 20}
	radius := 5.0

	c, _ := geometry.NewCircle(radius, center)
	bb := c.BoundingBox()

	if bb.MinX != center.X-radius {
		t.Errorf("MinX mismatch")
	}
	if bb.MinY != center.Y-radius {
		t.Errorf("MinY mismatch")
	}
	if bb.MaxX != center.X+radius {
		t.Errorf("MaxX mismatch")
	}
	if bb.MaxY != center.Y+radius {
		t.Errorf("MaxY mismatch")
	}
}
