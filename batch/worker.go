package batch

import "geomEngine/engine"

//Worker: long-living goroutine that repeatedly picks up tasks and processes them
func Worker(id int, tasks <-chan Task, results chan<- Result) { //worker can only read from tasks
	for task := range tasks {
		value, err := engine.Execute(task.Shape, task.Operation)

		results <- NewResult(task.ID, value, err) //send-only channel, worker can only write results
	}
}

// workers produce results
// processor consumes results
