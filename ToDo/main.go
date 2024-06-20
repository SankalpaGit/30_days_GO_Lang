// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	taskManager := NewTaskManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an option: (add, delete, view, exit)")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		switch strings.ToLower(input) {
		case "add":
			fmt.Print("Enter task description: ")
			if !scanner.Scan() {
				break
			}
			description := scanner.Text()

			fmt.Print("Enter deadline (YYYY-MM-DD HH:MM:SS): ")
			if !scanner.Scan() {
				break
			}
			deadlineStr := scanner.Text()
			deadline, err := time.Parse("2006-01-02 15:04:05", deadlineStr)
			if err != nil {
				fmt.Println("Invalid date format:", err)
				continue
			}
			taskManager.AddTask(description, deadline)

		case "delete":
			fmt.Print("Enter task ID to delete: ")
			if !scanner.Scan() {
				break
			}
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID format:", err)
				continue
			}
			taskManager.DeleteTask(id)

		case "view":
			taskManager.ViewTasks()

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
