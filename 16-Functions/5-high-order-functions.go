package main

import "fmt"

// greetFunction is a higher-order function that takes a string and returns another function.
func greetFunction(greeting string) func(string) string {
	// Return a new function that takes a string and returns a formatted greeting.
	return func(name string) string {
		return fmt.Sprintf("%s, %s!", greeting, name)
	}
}

func main() {
	// Create two functions that greet in different ways.
	englishGreeting := greetFunction("Hello")
	spanishGreeting := greetFunction("Hola")

	// Call the englishGreeting and spanishGreeting functions to get formatted greetings.
	fmt.Println(englishGreeting("Alice")) // Output: Hello, Alice!
	fmt.Println(spanishGreeting("Bob"))   // Output: Hola, Bob!
}
