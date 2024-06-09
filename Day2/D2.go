package main

import "fmt"

//Single value assignment
var day1 string = "sunday"

//Mutiple value assignment in Go Lang
var x, y int = 2, 3

// like js and some other languages we can use const also to declare variables
const PI = 3.14

func main() { //we need main function to run code
	fmt.Print(day1, "\n")
	print(PI, "\n")
	fmt.Print("The products is :", x*y)
}
