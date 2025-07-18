package internal

import (
	"encoding/json"
	"os"
)

var storageFile = "tasks.json"

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
}

func LoadTasks() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil // No file yet
		}
		return nil, err
	}
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
