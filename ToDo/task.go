// task.go
package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Description string
	Deadline    time.Time
}

func (t *Task) String() string {
	return fmt.Sprintf("ID: %d, Description: %s, Deadline: %s", t.ID, t.Description, t.Deadline.Format("2006-01-02 15:04:05"))
}
