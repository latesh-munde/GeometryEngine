package engine_test

import (
	"testing"

	"geomEngine/engine"
	"geomEngine/geometry"
)

func TestExecuteArea(t *testing.T) {
	circle, err := geometry.NewCircle(2, geometry.Point{X: 0, Y: 0})
	if err != nil {
		t.Fatalf("failed to create circle: %v", err)
	}

	result, err := engine.Execute(circle, engine.OpArea)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	area, ok := result.(float64)
	if !ok {
		t.Fatalf("expected float64 result, got %T", result)
	}

	if area <= 0 {
		t.Errorf("expected area > 0, got %v", area)
	}
}

func TestExecuteBoundingBox(t *testing.T) {
	rect, err := geometry.NewRectangle(10, 5.0, geometry.Point{X: 0, Y: 0})
	if err != nil {
		t.Fatalf("failed to create rectangle: %v", err)
	}

	result, err := engine.Execute(rect, engine.OpBoundingBox)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, ok := result.(geometry.BoundingBox)
	if !ok {
		t.Fatalf("expected BoundingBox result, got %T", result)
	}
}

func TestExecuteUnknownOperation(t *testing.T) {
	circle, _ := geometry.NewCircle(1, geometry.Point{X: 0, Y: 0})

	result, err := engine.Execute(circle, engine.Operation(999))
	if err == nil {
		t.Fatalf("expected error for unknown operation, got nil")
	}

	if result != nil {
		t.Fatalf("expected nil result for unknown operation, got %v", result)
	}
}
