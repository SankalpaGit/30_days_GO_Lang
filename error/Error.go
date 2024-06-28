package main

import (
	"errors"
	"fmt"
)

// Custom error
var ErrNegativeNumber = errors.New("We dont take negative numbers for sum in our program")

func sumPositiveIntegers(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, ErrNegativeNumber
	}
	return a + b, nil
}

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	result, err := sumPositiveIntegers(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Sum of %d and %d is %d\n", a, b, result)
	}
}
