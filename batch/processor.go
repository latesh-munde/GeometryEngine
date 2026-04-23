package batch

import (
	"errors"
	"sync"
)

type Processor struct {
	Workers int
}

// NewProcessor creates a batch processor with a fixed number of workers.
func NewProcessor(Workers int) *Processor {
	return &Processor{Workers: Workers}
}

// Process executes a batch of tasks and returns their results.
func (p *Processor) Process(tasks []Task) ([]Result, error) {
	if p.Workers <= 0 { //prevents silent deadlock
		return nil, errors.New("Workers must be > 0")
	}

	taskCh := make(chan Task)     //work input
	resultCh := make(chan Result) //work output

	// start workers
	for i := 0; i < p.Workers; i++ {
		go Worker(i, taskCh, resultCh)
	}

	// send tasks - separate goroutine
	go func() {
		for _, task := range tasks {
			taskCh <- task
		}
		close(taskCh) //signals workers: no more tasks
	}()

	results := make([]Result, 0, len(tasks))

	//collect results
	for i := 0; i < len(tasks); i++ {
		res := <-resultCh
		results = append(results, res)
	}
	// know exactly how many results to expect
	// no waitgroups needed
	// no channel close needed for results

	return results, nil
}

// taskCh <-chan Task - receive-only
// ProcessStream cannot send into taskCh
func (p *Processor) ProcessStream(taskCh <-chan Task) <-chan Result {
	resultCh := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(p.Workers)

	internalTaskCh := make(chan Task)

	// Start workers
	for i := 0; i < p.Workers; i++ {
		go func(id int) {
			defer wg.Done()
			Worker(id, internalTaskCh, resultCh)
		}(i)
	}

	go func() {
		for task := range taskCh {
			internalTaskCh <- task
		}
		close(internalTaskCh)
	}()

	// Close result channel after all workers finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
	// Returned <-chan Result - receive-only
	// Callers cannot send into the returned resultCh
}
