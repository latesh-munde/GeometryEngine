package dto

type TaskJSON struct {
	ID        int            `json: "id`
	ShapeType string         `json: "shape_type`
	Params    map[string]any `json: "params`
	Operation string         `json: "operation`
}

func NewTaskJSON(id int, shapeType string, params map[string]any, operation string) TaskJSON {
	return TaskJSON{
		ID:        id,
		ShapeType: shapeType,
		Params:    params,
		Operation: operation,
	}
}
