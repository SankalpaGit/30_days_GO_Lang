package main

import (
	"fmt"
)

func main() {
	var numOperands int
	var operation string
	var result float64

	// Step 1: Ask the user how many numbers to perform task
	fmt.Print("Enter the number of operands: ")
	fmt.Scan(&numOperands)

	if numOperands < 1 {
		fmt.Println("The number of operands must be at least 1.")
		return
	}

	// Create a slice to store the operands
	operands := make([]float64, numOperands)

	// Step 2: Collect the numbers from the user
	for i := 0; i < numOperands; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&operands[i])
	}

	// Step 3: Ask the user for the operation
	fmt.Print("Enter the operation (+, -, *, /): ")
	fmt.Scan(&operation)

	// Step 4: Perform the operation
	result = operands[0]
	for i := 1; i < numOperands; i++ {
		switch operation {
		case "+":
			result += operands[i]
		case "-":
			result -= operands[i]
		case "*":
			result *= operands[i]
		case "/":
			if operands[i] == 0 {
				fmt.Println("Error: Division by zero")
				return
			}
			result /= operands[i]
		default:
			fmt.Println("Error: Invalid operation")
			return
		}
	}

	// Step 5: Print the result
	fmt.Printf("The result is: %f\n", result)
}
