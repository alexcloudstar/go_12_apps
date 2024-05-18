package duties

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Duty struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func Init() []Duty {
	var tasks []Duty

	return tasks
}

func New(tasks *[]Duty) *[]Duty {
	fmt.Println("Please enter the task description:")
	var description string
    reader, err := bufio.NewReader(os.Stdin).ReadString('\n')

    if err != nil {
        fmt.Println("Something went wrong.")
    }

    description = strings.TrimSpace(reader)


	task := Duty{
		ID:          len(*tasks),
		Description: description,
		Completed:   false,
	}

    *tasks = append(*tasks, task)

    return tasks
}

func ToggleCompleted(tasks *[]Duty, task_id string) *[]Duty {
    for idx, el := range *tasks {
		task_id := parseId(task_id)
		if el.ID == int(task_id) {
            (*tasks)[idx].Completed = !el.Completed
		}
	}

    return tasks
}

func DeleteDuty(tasks *[]Duty, task_id string) *[]Duty {
    for idx, el := range *tasks {
		task_id := parseId(task_id)
		if el.ID == int(task_id) {
            *tasks = append((*tasks)[:idx], (*tasks)[idx+1:]...)
		}
	}

	return tasks
}

func parseId(id string) int64 {
	task_id, _ := strconv.ParseInt(id, 10, 64)

	return task_id
}
