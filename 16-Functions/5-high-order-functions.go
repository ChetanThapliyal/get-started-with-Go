// Calculating different parameters of Circle using high order function
package main

import "fmt"

// Function to calculate the area of a circle given its radius
func Area(r float64) float64 {
	return 3.14 * r * r
}

// Function to calculate the perimeter of a circle given its radius
func Perimeter(r float64) float64 {
	return 2 * 3.14 * r
}

// Function to calculate the diameter of a circle given its radius
func Diameter(r float64) float64 {
	return 2 * r
}

// Function to print the result of a calculation on a circle
// The calcFunction parameter is a function that takes the radius of a circle and returns a float64 result
func printResult(radius float64, calcFunction func(r float64) float64) {
	// If the calcFunction is nil (i.e., it wasn't provided or set to a function), print an error message and return
	if calcFunction == nil {
		fmt.Println("Invalid response")
		return
	}
	// Calculate the result using the provided calcFunction and print it, along with a thank you message
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you!")
}

// Function to return the appropriate function based on the user's query
func getFunction(query int) func(r float64) float64 {
	// Create a map that maps query numbers to the appropriate function
	query_to_func := map[int]func(r float64) float64{
		1: Area,
		2: Perimeter,
		3: Diameter,
	}
	// If the query number maps to a valid function, return that function
	if function, ok := query_to_func[query]; ok {
		return function
		// Otherwise, print an error message and return nil
	} else {
		fmt.Println("Invalid response")
		return nil
	}
}

func main() {
	var query int
	var radius float64
	// Ask the user for the radius of the circle and read it in
	fmt.Println("Enter the radius of circle: ")
	fmt.Scanf("%f", &radius)
	// Ask the user for the calculation they want to perform and read it in
	fmt.Println("Enter : \n 1-Area\n 2-Perimeter\n 3-Diameter")
	fmt.Scanf("%d", &query)
	// Get the appropriate function based on the user's query and print the result
	printResult(radius, getFunction(query))
}
