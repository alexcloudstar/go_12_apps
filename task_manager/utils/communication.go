package utils

import "fmt"

func ShowOptions() {
	fmt.Println("1. Add a new task")
	fmt.Println("2. List all tasks")
	fmt.Println("3. Mark a task as completed")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Exit")
}

func PrintSeparator() {
    fmt.Println("----------------------------------------------------")
}
