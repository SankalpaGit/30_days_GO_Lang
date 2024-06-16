package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("Original slice:", slice)

	slice[1] = 100
	slice[2] = 200

	fmt.Println("Modified slice:", slice)

	subSlice := slice[1:5] //start to end-1 is the rule for slicing

	// Print the new slice
	fmt.Println("Sub-slice:", subSlice)
}
