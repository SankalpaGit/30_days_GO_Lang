package main

import "net/http"

func main() {
	type PointerToBool = *bool
	type MapOfIntToInt = map[int]int // Alias for map with int keys and int values

	// Define a struct with various fields
	var exampleStruct struct {
		// Fields of different types:
		fieldString      string        // Regular string field
		fieldError       error         // Regular error interface field
		fieldInt         *int          // Pointer to an int (unnamed pointer type)
		fieldBoolPointer PointerToBool // Alias for pointer to bool (named pointer type)
		fieldIntMap      MapOfIntToInt // Alias for map[int]int (named map type)

		// Embedded field:
		http.Header // Embedding the http.Header type (a named map type)
	}

	// Assign values to the struct fields:
	exampleStruct.fieldString = "Go"                // Assign a string
	exampleStruct.fieldError = nil                  // Assign nil to the error interface
	exampleStruct.fieldInt = new(int)               // Allocate a new int and assign its address
	exampleStruct.fieldBoolPointer = new(bool)      // Allocate a new bool and assign its address
	exampleStruct.fieldIntMap = make(MapOfIntToInt) // Initialize the map using make
	exampleStruct.Header = http.Header{}            // Initialize the embedded http.Header map
}
