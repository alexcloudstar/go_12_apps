package duties

import (
	"fmt"
	"strconv"
)

type Duty struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func New(tasks []Duty) []Duty {
	fmt.Println("Please enter the task description:")
	var description string
	fmt.Scanln(&description)

	task := Duty{
		ID:          len(tasks),
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, task)

    return tasks
}

func ToggleCompleted (tasks []Duty, task_id string) []Duty {
    for idx, el := range tasks {
        task_id := parseId(task_id)
        if el.ID == int(task_id) {
            tasks[idx].Completed = true
        }
    }

    return tasks
}

func DeleteDuty(tasks []Duty, task_id string) []Duty {
    for idx, el := range tasks {
        task_id := parseId(task_id)
        if el.ID == int(task_id) {
            tasks = append(tasks[:idx], tasks[idx+1:]...)
        }
    }

    return tasks
}

func parseId(id string) int64 {
    task_id, _ := strconv.ParseInt(id, 10, 64)

    return task_id
}