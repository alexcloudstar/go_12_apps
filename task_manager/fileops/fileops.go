package fileops

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/alexcloudstar/go_12_apps/task_manager/duties"
)

const File_name = "tasks.json"
const File_mode = 0644

func Init(tasks *[]duties.Duty) []byte {
	content, err := os.ReadFile(File_name)

	if errors.Is(err, os.ErrNotExist) {
		file, _ := os.Create(File_name)

		file.Sync()
	}

	if len(content) > 0 {
		err = json.Unmarshal(content, &tasks)
	}

	return content
}

func Write(tasks *[]duties.Duty) error {
	json, _ := json.Marshal(&tasks)

	err := os.WriteFile(File_name, json, File_mode)

	return err
}
