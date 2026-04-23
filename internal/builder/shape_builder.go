package builder

import (
	"errors"
	"geomEngine/dto"
	"geomEngine/geometry"
)

var ErrUnknownShape = errors.New("unknown shape type")

func BuildShape(t dto.TaskJSON) (geometry.Shape, error) {
	switch t.ShapeType {
	case "circle":
		centerMap := t.Params["center"].(map[string]interface{})
		x := centerMap["x"].(float64)
		y := centerMap["y"].(float64)
		radius := t.Params["radius"].(float64)
		return geometry.NewCircle(radius, geometry.NewPoint(x, y))

	case "rectangle":
		originMap := t.Params["origin"].(map[string]interface{})
		x := originMap["x"].(float64)
		y := originMap["y"].(float64)
		width := t.Params["width"].(float64)
		height := t.Params["height"].(float64)
		return geometry.NewRectangle(width, height, geometry.NewPoint(x, y))

	case "line":
		startMap := t.Params["start"].(map[string]interface{})
		x1 := startMap["x"].(float64)
		y1 := startMap["y"].(float64)

		endMap := t.Params["end"].(map[string]interface{})
		x2 := endMap["x"].(float64)
		y2 := endMap["y"].(float64)
		return geometry.NewLine(geometry.NewPoint(x1, y1), geometry.NewPoint(x2, y2))

	case "polygon":
		rawPoints := t.Params["points"].([]interface{})
		points := make([]geometry.Point, 0, len(rawPoints))

		for _, p := range rawPoints {
			pt := p.(map[string]interface{})
			x := pt["x"].(float64)
			y := pt["y"].(float64)
			points = append(points, geometry.NewPoint(x, y))
		}
		return geometry.NewPolygon(points)
	}
	return nil, ErrUnknownShape
}
