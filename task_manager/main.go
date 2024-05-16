package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

const file = "tasks.json"

func main() {
	fmt.Println("Hi! Please choose an option:")

	printInitMsg()
	_, err := os.Create(file)

	if err != nil {
		fmt.Println("Could not create the tasks file")
		return
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
				ID:          0,
				Description: description,
				Completed:   false,
			}

			json, _ := json.Marshal(task)

			os.WriteFile(file, json, 0644)

		case 2:
			fmt.Println("Here are all the tasks:")
		case 3:
			fmt.Println("Please enter the task number you want to mark as completed:")
		case 4:
			fmt.Println("Please enter the task number you want to delete:")
		default:
			fmt.Println("Goodbye!")
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
