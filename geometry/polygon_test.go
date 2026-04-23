package geometry_test

import (
	"testing"

	"geomEngine/geometry"
)

func TestPolygonValidation(t *testing.T) {
	points := []geometry.Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	_, err := geometry.NewPolygon(points)
	if err != nil {
		t.Fatalf("expected valid polygon, got error: %v", err)
	}

	_, err = geometry.NewPolygon([]geometry.Point{
		{X: 0, Y: 0},
		{X: 1, Y: 1},
	})
	if err == nil {
		t.Fatalf("expected error for polygon with < 3 points")
	}
}

func TestPolygonArea(t *testing.T) {
	// Right triangle with area = 0.5
	points := []geometry.Point{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
	}

	p, _ := geometry.NewPolygon(points)
	area := p.Area()

	expected := 0.5
	if area != expected {
		t.Errorf("area mismatch: expected %v, got %v", expected, area)
	}
}

func TestPolygonPerimeter(t *testing.T) {
	points := []geometry.Point{
		{X: 0, Y: 0},
		{X: 3, Y: 0},
		{X: 3, Y: 4},
	}

	p, _ := geometry.NewPolygon(points)
	perimeter := p.Perimeter()

	expected := 12.0 // 3 + 4 + 5
	if perimeter != expected {
		t.Errorf("perimeter mismatch: expected %v, got %v", expected, perimeter)
	}
}

func TestPolygonBoundingBox(t *testing.T) {
	points := []geometry.Point{
		{X: -1, Y: 2},
		{X: 3, Y: 5},
		{X: 0, Y: -2},
	}

	p, _ := geometry.NewPolygon(points)
	bb := p.BoundingBox()

	if bb.MinX != -1 || bb.MaxX != 3 {
		t.Errorf("bounding box X incorrect")
	}
	if bb.MinY != -2 || bb.MaxY != 5 {
		t.Errorf("bounding box Y incorrect")
	}
}
