package generator

import (
	"encoding/json"
	"os"
)

func GenerateTasks(path string, totalTasks int) {
	file, _ := os.Create(path)
	defer file.Close()

	enc := json.NewEncoder(file)

	file.Write([]byte("["))

	for i := 0; i < totalTasks; i++ {
		task := randomTask(i)

		if i > 0 {
			file.Write([]byte(",")) //JSON arrays need commas
		}
		enc.Encode(task) //Encoder adds newline automatically
	}

	file.Write([]byte("]"))
}
