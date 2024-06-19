// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		case "delete":
			fmt.Print("Enter task ID to delete: ")

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
