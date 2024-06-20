// task_manager.go
package main

import (
	"fmt"
	"sync"
	"time"
)

type TaskManager struct {
	tasks  map[int]*Task
	mu     sync.Mutex
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

func (tm *TaskManager) AddTask(description string, deadline time.Time) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	task := &Task{
		ID:          tm.nextID,
		Description: description,
		Deadline:    deadline,
	}
	tm.tasks[tm.nextID] = task
	tm.nextID++
	fmt.Println("Task added:", task)
}

func (tm *TaskManager) DeleteTask(id int) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if _, exists := tm.tasks[id]; exists {
		delete(tm.tasks, id)
		fmt.Printf("Task with ID %d deleted.\n", id)
	} else {
		fmt.Printf("Task with ID %d not found.\n", id)
	}
}

func (tm *TaskManager) ViewTasks() {

}
