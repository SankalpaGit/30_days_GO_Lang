package main

import "fmt"

func updateElement(arr *[3]int, index int, newValue int) {
	if index >= 0 && index < len(arr) {
		(*arr)[index] = newValue
	} else {
		fmt.Println("Index out of range")
	}
}

func main() {
	array := [3]int{10, 20, 30}
	fmt.Println("Original array:", array)

	updateElement(&array, 0, 200)

	fmt.Println("Updated array:", array)
}
