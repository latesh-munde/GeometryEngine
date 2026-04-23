package runner

import (
	"geomEngine/batch"
	"geomEngine/engine"
)

func parseOperation(op string) engine.Operation {
	switch op {
	case "area":
		return engine.OpArea
	case "perimeter":
		return engine.OpPerimeter
	case "bbox":
		return engine.OpBoundingBox
	default:
		return engine.Operation(999)
	}
}

// Reads until channel closes
// Converts stream - slice
func readAllTasks(ch <-chan batch.Task) []batch.Task {
	var tasks []batch.Task
	for t := range ch {
		tasks = append(tasks, t)
	}
	return tasks
}
