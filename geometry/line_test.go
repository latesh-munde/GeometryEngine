package geometry_test

import (
	"math"
	"testing"

	"geomEngine/geometry"
)

const lineEpsilon = 0.0001

func TestLineValidation(t *testing.T) {
	start := geometry.Point{X: 0, Y: 0}
	end := geometry.Point{X: 1, Y: 1}

	_, err := geometry.NewLine(start, end)
	if err != nil {
		t.Fatalf("expected valid line, got error: %v", err)
	}

	_, err = geometry.NewLine(start, start)
	if err == nil {
		t.Fatalf("expected error for degenerate line, got nil")
	}
}

func TestLinePerimeter(t *testing.T) {
	start := geometry.Point{X: 0, Y: 0}
	end := geometry.Point{X: 3, Y: 4}

	l, _ := geometry.NewLine(start, end)

	expected := 5.0
	actual := l.Perimeter()

	if math.Abs(actual-expected) > lineEpsilon {
		t.Errorf("length mismatch: expected %v, got %v", expected, actual)
	}
}

func TestLineBoundingBox(t *testing.T) {
	start := geometry.Point{X: 5, Y: 2}
	end := geometry.Point{X: 1, Y: 8}

	l, _ := geometry.NewLine(start, end)
	bb := l.BoundingBox()

	if bb.MinX != 1 || bb.MaxX != 5 {
		t.Errorf("bounding box X incorrect")
	}
	if bb.MinY != 2 || bb.MaxY != 8 {
		t.Errorf("bounding box Y incorrect")
	}
}
