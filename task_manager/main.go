package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alexcloudstar/go_12_apps/task_manager/duties"
	"github.com/alexcloudstar/go_12_apps/task_manager/utils"
)

const file_name = "tasks.json"
const osPerm = os.O_CREATE | os.O_WRONLY
const file_mode = 0644

func main() {
	fmt.Println("Hi! Please choose an option:")
    tasks := duties.Init()

    utils.ShowOptions()
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
            tasks := duties.New(tasks)

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

             tasks := duties.ToggleCompleted(tasks, task_id)

            json, _ := json.Marshal(&tasks)

            err := os.WriteFile(file_name, json, file_mode)

            if err != nil {
                fmt.Println("could not write to the tasks file")
                fmt.Println(err)
            }
			
		case 4:
			fmt.Println("Please enter the task number you want to delete:")
			var task_id string
			fmt.Scanln(&task_id)

            tasks := duties.DeleteDuty(tasks, task_id)
            json, _ := json.Marshal(&tasks)

            err := os.WriteFile(file_name, json, file_mode)

            if err != nil {
                fmt.Println("Could not write to the tasks file")
                fmt.Println(err)
            }
		case 5:
            fmt.Println("Goodbye!")
			return
		}

        fmt.Println("\n")
        utils.ShowOptions()
	}
}

