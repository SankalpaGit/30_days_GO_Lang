package main

import "fmt"

func main() {
	var marks = [5]int{50, 55, 60, 80, 65} //If we want to use var for variable declaration
	marking := [3]int{40, 55, 60}          //If we dont prefer to use var for variable declaration

	/* remeber if we use var we would used = for the assignment of value
	where in case not using we have to use := for the assignment of the value*/

	fmt.Println(marks)
	fmt.Println(marking)

	//Using the for loop in go lang and if statement

	for i := 0; i < len(marks); i++ { //loop start from 0 to length of marks with auto increment
		if marks[i] >= 60 { //checking the condition if marks is greter or equal to 60
			fmt.Println(marks[i])
		}
	}

	//checking the all the value in marking array now
	passingCriteria := true
	for i := 0; i < len(marking); i++ {
		if marking[i] <= 50 {
			passingCriteria = false
			break
		}
	}

	if passingCriteria {
		fmt.Println("You have an average above 50")
	} else {
		fmt.Println("Work hard for geting avg 50")
	}
}
