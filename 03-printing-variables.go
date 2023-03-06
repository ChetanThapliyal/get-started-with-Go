package main

import "fmt"

func main() {

	//fmt.Printf : to print a formatted string to the console, which can include variables
	//format specifiers:
	//`%v`: Prints the value in a default format.
	//`%d`: Prints an integer value in decimal format.
	//`%f`: Prints a floating-point value in decimal format.
	//`%s`: Prints a string value.
	//`%t`: Prints a boolean value.
	//`%c`: Prints a character value.

	name := "John Doe"
	age := 26
	fmt.Printf("My name is %s and I am %d years old\n", name, age)

	//fmt.Println: adds a newline character at the end
	var user string = "Harry"
	fmt.Println(name)
	fmt.Println(user)

}

