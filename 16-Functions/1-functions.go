package main

import "fmt"

//Function declaration syntax:
//func function_name(Parameter-list)(Return_type){function body...}

// Function declaration with two parameters
func area(length, width int) int {

	Ar := length * width
	return Ar
}

func main() {

	// Function calling with two arguments
	area := area(12, 10)
	fmt.Println("Area of rectangle is : ", area)
}
