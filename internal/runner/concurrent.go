package runner

import (
	"encoding/json"
	"geomEngine/batch"
	"geomEngine/dto"
	"geomEngine/internal/builder"
	"os"
	"runtime"
	"time"
)

func RunConcurrent(taskFile string) time.Duration {
	file, _ := os.Open(taskFile)
	defer file.Close()

	dec := json.NewDecoder(file)
	dec.Token() // [

	// NumCPU returns the number of logical CPUs usable by the current process.
	processor := batch.NewProcessor(runtime.NumCPU())
	taskCh := make(chan batch.Task) // taskCh feeds tasks to workers

	// Producer goroutine
	go func() {
		defer close(taskCh)
		for dec.More() {
			var tj dto.TaskJSON
			//Decode reads the next JSON-encoded value from its input
			// and stores it in the value pointed to by tj.
			dec.Decode(&tj) //Converts DTO → runtime task

			shape, err := builder.BuildShape(tj)
			if err != nil {
				continue
			}

			taskCh <- batch.NewTask(tj.ID, shape, parseOperation(tj.Operation))
		}
	}()

	start := time.Now()
	results, _ := processor.Process(readAllTasks(taskCh))
	// readAllTasks collects tasks
	// processor.Process distributes tasks to workers
	_ = results //Results are ignored

	return time.Since(start)
}
