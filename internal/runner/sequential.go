package runner

import (
	"encoding/json"
	"geomEngine/dto"
	"geomEngine/engine"
	"geomEngine/internal/builder"
	"os"
	"time"
)

func RunSequential(taskFile string) time.Duration {
	file, _ := os.Open(taskFile)
	defer file.Close()

	dec := json.NewDecoder(file)
	dec.Token() // [ // skips

	start := time.Now()

	//More reports whether there is another element in the current array, file, etc
	for dec.More() {
		var tj dto.TaskJSON
		dec.Decode(&tj) // value read from file is saved in tj

		shape, err := builder.BuildShape(tj)
		if err != nil {
			continue
		}

		op := parseOperation(tj.Operation)

		engine.Execute(shape, op)
	}

	return time.Since(start)
}
