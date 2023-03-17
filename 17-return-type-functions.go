package main

import "fmt"

// A function that takes two integers and returns their sum as a single value.
func add(a, b int) int {
	return a + b
}

// A function that takes two integers and returns their sum and product as multiple values.
func addAndMultiply(a, b int) (int, int) {
	return a + b, a * b
}

// A function that takes two integers and returns their difference and quotient as named return values.
func subtractAndDivide(a, b int) (diff, quot int) {
	diff = a - b
	quot = a / b
	return // returns both named values
}

// A variadic function that takes any number of integers and returns their sum.
func sumNum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	// Calling the single return value function.
	fmt.Println(add(5, 10)) //Output: 15

	// Calling the multiple return value function.
	sum, product := addAndMultiply(5, 10)
	fmt.Println(sum, product) //Output: 15,50

	// Calling the named return value function.
	diff, quot := subtractAndDivide(5, 2)
	fmt.Println(diff, quot) //Output: 3,2

	// Calling the variadic function.
	fmt.Println(sumNum(1, 2, 3, 4, 5)) //Output: 15
}

