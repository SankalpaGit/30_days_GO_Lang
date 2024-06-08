package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* Different ways for running the Go file
	   - Run through the terminal: go run file.go
	   - Run through the Run task in IDE or editor
	   - By converting the file to binary format
	*/
	fmt.Println("Hello world! I am a GoLang Developer")

	// Taking input for the first number
	var input1 string
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&input1)

	// Converting the input to an integer
	num1, err := strconv.Atoi(input1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Taking input for the second number
	var input2 string
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&input2)

	// Converting the input to an integer
	num2, err := strconv.Atoi(input2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Performing arithmetic operations
	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := float64(num1) / float64(num2)

	// Printing the results
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
}
