package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

const file_name = "tasks.json"
const osPerm = os.O_CREATE | os.O_WRONLY
const file_mode = 0644

func main() {
	fmt.Println("Hi! Please choose an option:")
	var tasks []Task

	printInitMsg()
	file, err := os.OpenFile(file_name, osPerm, file_mode)

	defer file.Close()

	if err != nil {
		fmt.Println("Could not create the tasks file")
		return
	}

	content, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Could not read the tasks file")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(content) > 0 {
		err = json.Unmarshal(content, &tasks)
	}

	for {
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Please enter the task description:")
			var description string
			fmt.Scanln(&description)

			task := Task{
				ID:          len(tasks),
				Description: description,
				Completed:   false,
			}

			tasks = append(tasks, task)

			json, _ := json.Marshal(&tasks)

			err := os.WriteFile(file_name, json, file_mode)

			if err != nil {
				fmt.Println("Could not write to the tasks file")
				fmt.Println(err)
			}
		case 2:
			fmt.Println("Here are all the tasks:")
			fmt.Println(string(content))
		case 3:
			fmt.Println("Please enter the task number you want to mark as completed:")
			var task_id string
			fmt.Scanln(&task_id)

			for idx, el := range tasks {
				task_id_int, _ := strconv.ParseInt(task_id, 10, 64)
				if el.ID == int(task_id_int) {
					tasks[idx].Completed = true
					json, _ := json.Marshal(&tasks)

					err := os.WriteFile(file_name, json, file_mode)

					if err != nil {
						fmt.Println("Could not write to the tasks file")
						fmt.Println(err)
					}

				}
			}
		case 4:
			fmt.Println("Please enter the task number you want to delete:")
			var task_id string
			fmt.Scanln(&task_id)

			for idx, el := range tasks {
				task_id_int, _ := strconv.ParseInt(task_id, 10, 64)
				if el.ID == int(task_id_int) {
					tasks = append(tasks[:idx], tasks[idx+1:]...)
					json, _ := json.Marshal(&tasks)

					err := os.WriteFile(file_name, json, file_mode)

					if err != nil {
						fmt.Println("Could not write to the tasks file")
						fmt.Println(err)
					}

				}
			}
		case 5:
			return
		}
	}
}

func printInitMsg() {
	fmt.Println("1. Add a new task")
	fmt.Println("2. List all tasks")
	fmt.Println("3. Mark a task as completed")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Exit")
}
