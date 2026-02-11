package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type TaskStatus int

const (
	StatusTodo TaskStatus = iota
	StatusInProgress
	StatusDone
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskManager struct {
	tasks   []Task
	Next_Id int
}

func (s TaskStatus) String() string {
	switch s {
	case StatusTodo:
		return "Todo"
	case StatusInProgress:
		return "In-Progress"
	case StatusDone:
		return "Done"
	default:
		return "unknown"
	}
}

func ParseStatus(input string) TaskStatus {
	switch strings.ToLower(strings.TrimSpace(input)) {
	case "todostatus", "todo":
		return StatusTodo
	case "statusinprogress", "in-progress", "inprogress":
		return StatusInProgress
	case "statusdone", "done", "Done":
		return StatusDone
	default:
		return 0
	}
}

func (tm *TaskManager) LoadTask() error {
	data, err := os.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("Failed to read tasks: %w ", err)
	}

	if len(data) != 0 {
		err = json.Unmarshal(data, &tm.tasks)
		if err != nil {
			return fmt.Errorf("Failed to parse tasks: %w ", err)
		}
	}
	return nil
}

func handleAddTask(tm *TaskManager) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task add \"task description\"")
		return
	}

	task := os.Args[2]
	tm.AddTask(task)

}

func handleUpdateTask(tm *TaskManager) {
	if len(os.Args) < 4 {
		fmt.Println("Usage: task update <id> \"updated task description\"")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid id")
		return
	}
	updatedTask := os.Args[3]
	tm.UpdateTask(id, updatedTask)
}

func handleDeleteTask(tm *TaskManager) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: task delete <id>")
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid id")
		return
	}
	fmt.Println(id, "this is the id ")
	tm.DeleteTask(id)
}

func handleUpdateStatus(tm *TaskManager) {
	if len(os.Args) < 4 {
		fmt.Println("Usage: task mark in-progress | done <id>")
	}

	id, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Print("invalid id ")
	}
	status := os.Args[2]
	taskStatus := ParseStatus(status)
	tm.UpdateStatus(id, taskStatus)
}

func handleListTask(tm *TaskManager, val string) {
	if val == "all" {
		if len(os.Args) < 2 {
			fmt.Println("Usage: task list")
			return
		}

		tm.ListAllTask(val)
	} else if val == "filter" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: task list status")
			return
		}
		status := os.Args[2]
		tm.ListAllTask(status)
	}
}

func (tm *TaskManager) AddTask(task string) {
	if len(task) == 0 {
		return
	}
	if len(tm.tasks) > 0 {
		tm.Next_Id = tm.tasks[len(tm.tasks)-1].ID + 1
	} else {
		tm.Next_Id = 1
	}
	newTask := Task{
		ID:          tm.Next_Id,
		Description: task,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tm.tasks = append(tm.tasks, newTask)

	if err := tm.save(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to save task: %v\n", err)
		return
	}
	fmt.Println("Added New Task successfully")
}

func (tm *TaskManager) UpdateTask(id int, newTask string) {
	if id <= 0 || newTask == "" {
		fmt.Println("invalid input")
		return
	}
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Description = newTask
			tm.tasks[i].UpdatedAt = time.Now()
		}
	}
	if err := tm.save(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to save task: %v\n", err)
		return
	}
	fmt.Println("Updated Successfully")
}

func (tm *TaskManager) DeleteTask(id int) {
	if id <= 0 {
		fmt.Print("Invalid input")
		return
	}

	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			if err := tm.save(); err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to save task: %v\n", err)
				return
			}
			fmt.Println("Deleted Successfully")
			return
		}
	}
	fmt.Println("Task not found")
}

func (tm *TaskManager) UpdateStatus(id int, status TaskStatus) {
	if id <= 0 {
		fmt.Println("invalid input")
		return
	}

	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Status = status
			tm.tasks[i].UpdatedAt = time.Now()
		}
	}
	if err := tm.save(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to save task: %v\n", err)
		return
	}
	fmt.Println("Status updated successfully")
}

func (tm *TaskManager) ListAllTask(filter string) {
	if len(tm.tasks) == 0 {
		fmt.Print("No task found!")
		return
	}

	for _, val := range tm.tasks {
		switch filter {
		case "all":
			fmt.Printf("%d. %s | %s\n", val.ID, val.Description, val.Status)
		case "done":
			if val.Status == StatusDone {
				fmt.Printf("%d. %s | %s\n", val.ID, val.Description, val.Status)
			}
		case "todo":
			if val.Status == StatusTodo {
				fmt.Printf("%d. %s | %s\n", val.ID, val.Description, val.Status)
			}
		case "in-progress":
			if val.Status == StatusInProgress {
				fmt.Printf("%d. %s | %s\n", val.ID, val.Description, val.Status)
			}
		}
	}
}

func (tm *TaskManager) save() error {
	data, err := json.MarshalIndent(tm.tasks, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to serialize tasks: %w ", err)
	}
	err = os.WriteFile("data.json", data, 0644)
	if err != nil {
		return fmt.Errorf("Failed to save tasks: %w ", err)
	}
	return nil
}

func main() {

	if _, err := os.Stat("data.json"); os.IsNotExist(err) {
		file, err := os.Create("data.json")
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	var tm TaskManager
	tm.LoadTask()

	if len(os.Args) < 2 {
		fmt.Println("Usage: task <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAddTask(&tm)
	case "update":
		handleUpdateTask(&tm)
	case "delete":
		handleDeleteTask(&tm)
	case "mark":
		handleUpdateStatus(&tm)
	case "list":
		if len(os.Args) < 3 {
			handleListTask(&tm, "all")
		} else {
			handleListTask(&tm, "filter")
		}
	}

}
