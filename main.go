package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type TasksTraker struct {
	Tasks map[int]Task
}

type Task struct {
	Description string
	Status      string
	CreateAt    time.Time
	UpdateAt    time.Time
}

func (t *TasksTraker) add_task(message string) {
	id := rand.Intn(50)
	t.Tasks[id] = Task{
		Description: message,
		Status:      "todo",
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
}

func (t *TasksTraker) list_task() {
	for id, task := range t.Tasks {
		messages := fmt.Sprintf("Task id: %d, description: %s, status: %s, crear_at: %s, update_at: %s", id, task.Description, task.Status, task.CreateAt, task.UpdateAt)
		fmt.Println(messages)
	}
}

func (t *TasksTraker) find_one_task(id int) (*Task, error) {
	if task, ok := t.Tasks[id]; ok {
		return &task, nil
	}
	return nil, fmt.Errorf("task with id %d not found", id)
}

func (t *TasksTraker) find_filter_statu(statu string) (*Task, error) {

	for _, task := range t.Tasks {

		if task.Status == "done" {
			return &task, nil
		}
		if task.Status == "in-progress" {
			return &task, nil
		}
		if task.Status == "todo" {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task with id %s not found", statu)
}

func (t *TasksTraker) update_task(id int, changes string) (*Task, error) {
	task, err := t.find_one_task(id)
	if err != nil {
		fmt.Printf("Error finding task: %v\n", err)
		return nil, err
	}

	updatedTask := Task{
		Description: changes,
		Status:      task.Status,
		CreateAt:    task.CreateAt,
		UpdateAt:    time.Now(),
	}
	t.Tasks[id] = updatedTask
	return &updatedTask, nil
}

func (t *TasksTraker) delete_task(id int) (int, error) {
	_, err := t.find_one_task(id)
	if err != nil {
		return 0, err
	}
	delete(t.Tasks, id)
	return id, nil
}

func (t *TasksTraker) mark_task(id int, statu string) {
	task, _ := t.find_one_task(id)

	updata_statu := Task{
		Description: task.Description,
		Status:      statu,
		CreateAt:    task.CreateAt,
		UpdateAt:    time.Now(),
	}

	t.Tasks[id] = updata_statu
}

func main() {
	var flag string
	if len(os.Args) >= 2 {
		flag = os.Args[1]
	} else {
		flag = ""
	}

	var tasks TasksTraker
	data, err := os.ReadFile("Task.json")

	if err != nil {
		tasks = TasksTraker{
			Tasks: make(map[int]Task),
		}
	} else {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Printf("Error al leer el JSON: %v\n", err)
			return
		}
	}

	switch flag {
	case "-add":
		tasks.add_task(os.Args[2])
	case "-list":
		if len(os.Args) < 3 {
			tasks.list_task()
		} else {
			id := os.Args[2]
			task, err := tasks.find_filter_statu(id)
			if err != nil {
				fmt.Printf("Error finding task: %v\n", err)
				return
			}
			fmt.Printf("Task: %+v\n", *task)
		}
	case "-update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error converting id: %v\n", err)
			return
		}
		task, err := tasks.update_task(id, os.Args[3])
		if err != nil {
			fmt.Printf("Error finding task: %v\n", err)
			return
		}
		fmt.Printf("Task: %+v\n", *task)
	case "-delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error converting id: %v\n", err)
			return
		}
		id, err = tasks.delete_task(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	case "-mark-in-progress":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error converting id: %v\n", err)
			return
		}
		tasks.mark_task(id, "in-progress")
	case "-mark-done":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error converting id: %v\n", err)
			return
		}
		tasks.mark_task(id, "done")
	default:
		fmt.Println("invalid option")
		return
	}

	json_data, err := json.Marshal(&tasks)

	if err != nil {
		fmt.Printf("Error al codificar el JSON: %v\n", err)
		return
	}

	err = os.WriteFile("Task.json", json_data, 0644)

	if err != nil {
		fmt.Printf("Error al crear archivo: %v\n", err)
		return
	}
}
