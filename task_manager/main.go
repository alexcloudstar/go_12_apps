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

const file_name = "tasks.json"
const osPerm = os.O_APPEND | os.O_CREATE | os.O_WRONLY
const file_mode = 0644

func main() {
	fmt.Println("Hi! Please choose an option:")
	var tasks []Task

	printInitMsg()
	file, err := os.OpenFile(file_name, osPerm, file_mode)

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
				ID:          len(tasks),
				Description: description,
				Completed:   false,
			}

			tasks = append(tasks, task)

			json, _ := json.Marshal(tasks)

			_, err := file.Write(json)

			if err != nil {
				fmt.Println("Could not write to the tasks file")
				fmt.Println(err)
			}
		case 2:
			fmt.Println("Here are all the tasks:")
		case 3:
			fmt.Println("Please enter the task number you want to mark as completed:")
		case 4:
			fmt.Println("Please enter the task number you want to delete:")
			/*
			   FIX: this
			   default:
			   fmt.Println("Goodbye!")
			   return
			*/
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
