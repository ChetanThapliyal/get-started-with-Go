package main

import "fmt"

func main() {

	// Declaring a map with string keys and integer values
	var x map[string]int
	fmt.Println(x) // Output: map[]

	// Initializing an empty map with make function
	y := make(map[string]int)
	fmt.Println(y) // Output: map[]

	// Declaring and initializing a map with string keys and string values
	n := map[string]string{"foo": "bar", "baz": "qux"}
	fmt.Println(n) // Output:

	// Initializing a map with string keys and integer values
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m) // Output: map[one:1 two:2]

	//length of a map
	fmt.Println(len(m)) // Output: 2

	// Accessing an element in a map
	value := m["one"]
	fmt.Println(value) // Output: 1

	// Modifying/Updating key-value pair
	m["one"] = 10
	fmt.Println(m) // Output: map[one:10 two:2]

	// Adding a new element to a map
	m["three"] = 3
	fmt.Println(m) // Output: map[one:10 three:3 two:2]

	//deleting key-value pair
	delete(m, "two")
	fmt.Println(m) // Output: map[one:10 three:3]

	// Checking if a key exists in a map
	value, exists := m["one"]
	if exists {
		fmt.Println(value) // Output: 10
	} else {
		fmt.Println("Key does not exist")
	}

	// Iterating over a map
	for key, value := range m {
		fmt.Println(key, ":", value)  // Output: one : 10 three : 3
	}
}

