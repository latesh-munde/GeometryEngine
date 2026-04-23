package generator

import (
	"geomEngine/dto"
	"math/rand"
)

func randomTask(id int) dto.TaskJSON {
	switch rand.Intn(4) {
	case 0: // circle
		return dto.NewTaskJSON(
			id,
			"circle",
			map[string]any{
				"radius": randFloat(1, 50),
				"center": map[string]float64{
					"x": randFloat(-100, 100),
					"y": randFloat(-100, 100),
				},
			},
			randomOperation(),
		)
	case 1:
		return dto.NewTaskJSON(
			id,
			"rectangle",
			map[string]any{
				"width":  randFloat(1, 50),
				"height": randFloat(1, 50),
				"origin": map[string]float64{
					"x": randFloat(-100, 100),
					"y": randFloat(-100, 100),
				},
			},
			randomOperation(),
		)

	case 2:
		x1, y1 := randFloat(-100, 100), randFloat(-100, 100)
		x2, y2 := x1+randFloat(1, 50), y1+randFloat(1, 50)
		return dto.NewTaskJSON(
			id,
			"line",
			map[string]any{
				"start": map[string]float64{"x": x1, "y": y1},
				"end":   map[string]float64{"x": x2, "y": y2},
			},
			randomOperation(),
		)

	default:
		return dto.NewTaskJSON(
			id,
			"polygon",
			map[string]any{
				"points": []map[string]float64{
					{"x": randFloat(0, 10), "y": randFloat(0, 10)},
					{"x": randFloat(10, 20), "y": randFloat(0, 10)},
					{"x": randFloat(5, 15), "y": randFloat(10, 20)},
				},
			},
			randomOperation(),
		)

	}
}

func randomOperation() string {
	ops := []string{"area", "perimeter", "bbox"}
	return ops[rand.Intn(len(ops))]
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(min-max)
}
