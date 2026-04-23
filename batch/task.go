package batch

import (
	"geomEngine/engine"
	"geomEngine/geometry"
)

type Task struct {
	ID        int
	Shape     geometry.Shape
	Operation engine.Operation
}

// func NewTask(id int, shape geometry.Shape, operation engine.Operation) (*Task, error) {
// 	if id <= 0 || shape == nil || operation < 0 || operation > 2 {
// 		return &Task{}, errors.New("invalid arguments")
// 	}

// 	return &Task{
// 		ID:        id,
// 		Shape:     shape,
// 		Operation: operation,
// 	}, nil
// }

func NewTask(id int, shape geometry.Shape, operation engine.Operation) Task {
	return Task{
		ID:        id,
		Shape:     shape,
		Operation: operation,
	}
}
