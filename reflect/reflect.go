package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [3]int{10, 20, 30}        //declaring array
	fmt.Println("real array: ", array) //prints array

	modified := reflect.ValueOf(&array).Elem()
	modified.Index(1).SetInt(5) //changing the value of index 1 to 5
	fmt.Println("first modified: ", modified)

	checkType := modified.Index(1).Type()
	fmt.Println(checkType) //checking the type
	//we can find the lenght also with Len()

	newArray := []int{1, 2, 3}
	fmt.Println("Original second array:", newArray)

	// Creating a reflect.Value of the slice
	sliceVal := reflect.ValueOf(newArray)

	// Appending a new value to the slice using reflection
	newSliceVal := reflect.Append(sliceVal, reflect.ValueOf(4))

	// Convert the reflect.Value back to a slice interface and cast to the original type
	resultSlice := newSliceVal.Interface().([]int)

	// Printing the slices
	fmt.Println("Appended Slice:", resultSlice)

}
