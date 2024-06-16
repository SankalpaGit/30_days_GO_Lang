package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original array:", array)

	ptr1 := &array[0]
	ptr2 := &array[1]

	*ptr1 = 100
	*ptr2 = 200

	fmt.Println("Modified array:", array)
}
