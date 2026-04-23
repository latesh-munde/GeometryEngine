package batch_test

import (
	"geomEngine/batch"
	"geomEngine/engine"
	"geomEngine/geometry"
	"testing"
)

func TestBatchProcessSuccess(t *testing.T) {
	processor := batch.NewProcessor(3)

	c1, _ := geometry.NewCircle(2, geometry.NewPoint(0, 0))
	c2, _ := geometry.NewCircle(3, geometry.NewPoint(1, 1))

	tasks := []batch.Task{
		batch.NewTask(1, c1, engine.OpArea),
		batch.NewTask(2, c2, engine.OpPerimeter),
	}

	results, err := processor.Process(tasks)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != len(tasks) {
		t.Fatalf("expected %d results, got %d", len(tasks), len(results))
	}

	seen := make(map[int]bool)
	for _, res := range results {
		if res.Error != nil {
			t.Fatalf("unexpected task error: %v", res.Error)
		}
		seen[res.TaskID] = true
	}

	for _, task := range tasks {
		if !seen[task.ID] {
			t.Errorf("missing result for task %d", task.ID)
		}
	}
}

func TestBatchProcessWithInvalidOperation(t *testing.T) {
	processor := batch.NewProcessor(2)

	circle, _ := geometry.NewCircle(1, geometry.NewPoint(0, 0))
	tasks := []batch.Task{
		batch.NewTask(1, circle, engine.Operation(999)),
	}

	results, err := processor.Process(tasks)
	if err != nil {
		t.Fatalf("unexpected processor error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Error == nil {
		t.Fatalf("expected error for invalid operation, got nil")
	}
}
