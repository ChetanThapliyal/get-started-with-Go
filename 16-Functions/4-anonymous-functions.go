package main

import "fmt"

func main() {
	// Define and call an anonymous function
	func() {
		fmt.Println("Hello, anonymous function!")
	}()

	// Assign an anonymous function to a variable and call it
	add := func(a, b int) int {
		return a + b
	}

	sum := add(3, 4)
	fmt.Println(sum)
	fmt.Printf("%T \n", add)
}
