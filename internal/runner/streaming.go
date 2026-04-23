package runner

import (
	"encoding/json"
	"geomEngine/batch"
	"geomEngine/dto"
	"geomEngine/geometry"
	"geomEngine/internal/builder"
	"os"
	"runtime"
	"time"
)

func RunConcurrentStreaming(taskFile string) time.Duration {
	file, _ := os.Open(taskFile)
	defer file.Close()

	dec := json.NewDecoder(file)
	dec.Token()

	taskCh := make(chan batch.Task)

	go func() {
		defer close(taskCh)
		for dec.More() {
			var tj dto.TaskJSON
			dec.Decode(&tj) //Converts DTO → runtime task

			shape, err := builder.BuildShape(tj)
			if err != nil {
				continue
			}

			taskCh <- batch.NewTask(tj.ID, shape, parseOperation(tj.Operation))
		}
	}()

	processor := batch.NewProcessor(runtime.NumCPU())

	start := time.Now()

	results := processor.ProcessStream(taskCh)

	for res := range results {
		if res.Error != nil {
			continue
		}

		switch res.Value.(type) {
		case float64, geometry.BoundingBox:
		}

	}

	return time.Since(start)
}
