package main

import (
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func main() {
	fmt.Println("Hi! Please choose an option:")

	printInitMsg()
	os.Create("tasks.json")

	for {
		var option int
		fmt.Scanln(&option)

		if option < 1 || option > 5 {
			fmt.Println("Invalid option. Please choose a valid option.")
			printInitMsg()
			continue
		}

		switch option {
		case 1:
			fmt.Println("Please enter the task description:")
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
