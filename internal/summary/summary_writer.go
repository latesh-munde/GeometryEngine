package summary

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
)

func WriteSummary(path string, totalTasks int, seq, con time.Duration) {
	file, _ := os.Create(path)
	defer file.Close()

	summary := map[string]any{
		"tasks":                 totalTasks,
		"workers":               runtime.NumCPU(),
		"sequential_ms":         seq.Milliseconds(),
		"concurrent_ms":         con.Milliseconds(),
		"sequential_throughput": float64(totalTasks) / seq.Seconds(),
		"concurrent_throughput": float64(totalTasks) / con.Seconds(),
	}

	json.NewEncoder(file).Encode(summary)

	fmt.Println("Summary written to results/summary.json")
}
