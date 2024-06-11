package main

import "fmt"

func main() {
loop:
	for i := 0; i <= 10; i += 2 { // The loop runs until it reaches 10 with 2 increments
		switch {
		case i == 4:
			continue // Ignore the value 4 and skip to the next iteration
		case i == 8:
			break loop // Break the loop entirely when the value is 8
		}
		fmt.Println(i) // Print the value
	}
}
