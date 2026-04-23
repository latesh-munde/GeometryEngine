package main

import (
	"geomEngine/internal/generator"
	"geomEngine/internal/runner"
	"geomEngine/internal/summary"
)

const (
	TaskFile    = "results/task.json"
	SummaryFile = "results/summary.json"
	TotalTasks  = 100000
)

func main() {
	generator.GenerateTasks(TaskFile, TotalTasks)

	seqDuration := runner.RunSequential(TaskFile)

	// conDuration := runner.RunConcurrent(TaskFile)
	conDuration := runner.RunConcurrentStreaming(TaskFile)

	summary.WriteSummary(SummaryFile, TotalTasks, seqDuration, conDuration)
}
