package main

import "fmt"

func main() {
	fmt.Println("Hi! Please choose an option:")
	fmt.Println("1. Add a new task")
	fmt.Println("2. List all tasks")
	fmt.Println("3. Mark a task as completed")
	fmt.Println("4. Delete a task")

	for {
		var option int
		fmt.Scanln(&option)

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
