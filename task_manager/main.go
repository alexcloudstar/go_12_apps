package main

import (
	"fmt"

	"github.com/alexcloudstar/go_12_apps/task_manager/duties"
	"github.com/alexcloudstar/go_12_apps/task_manager/fileops"
	"github.com/alexcloudstar/go_12_apps/task_manager/utils"
)

const file_name = "tasks.json"
const file_mode = 0644

func main() {
	fmt.Println("Hi! Please choose an option:")
	tasks := duties.Init()

	utils.ShowOptions()

	content := fileops.Init(&tasks)

	for {
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			tasks := duties.New(&tasks)

			err := fileops.Write(tasks)

			if err != nil {
				fmt.Println("Could not write to the tasks file")
			}

		case 2:
			fmt.Println("Here are all the tasks:")
			fmt.Println(string(content))
		case 3:
			fmt.Println("Please enter the task number you want to mark as completed:")
			var task_id string
			fmt.Scanln(&task_id)

            tasks := duties.ToggleCompleted(&tasks, task_id)

			err := fileops.Write(tasks)

			if err != nil {
				fmt.Println("could not write to the tasks file")
				fmt.Println(err)
			}

		case 4:
			fmt.Println("Please enter the task number you want to delete:")
			var task_id string
			fmt.Scanln(&task_id)

			tasks := duties.DeleteDuty(&tasks, task_id)

			err := fileops.Write(tasks)

			if err != nil {
				fmt.Println("Could not write to the tasks file")
				fmt.Println(err)
			}
		case 5:
			fmt.Println("Goodbye!")
			return
		}
	}
}
