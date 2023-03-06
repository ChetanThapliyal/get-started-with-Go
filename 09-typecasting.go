package main

import "fmt"

func main() {

	//--------------------------------------------
	// Method 1: Using type conversion functions |
	//--------------------------------------------
	
	// Convert an int to a float64
	x := 10         //var x int =10
	y := float64(x) //var y float64 = float64(x)
	fmt.Println(y)

	// Convert a float64 to an int
	z := 10.5
	w := int(z)
	fmt.Println(w)

	// Convert a string to a byte slice
	s := "hello"
	b := []byte(s)
	fmt.Println(b)

	// Convert a byte slice to a string
	b2 := []byte{'h', 'e', 'l', 'l', 'o'}
	s2 := string(b2)
	fmt.Println(s2)

	//-------------------------------------------
	// Method 2: Using strconv package          |
	//-------------------------------------------

	// Converting a string to an integer
	numStr := "42"
	num, _ := strconv.Atoi(numStr)
	fmt.Printf("String %s converted to integer %d\n", numStr, num)

	// Converting an integer to a string
	numInt := 12345
	numStr = strconv.Itoa(numInt)
	fmt.Printf("Integer %d converted to string %s\n", numInt, numStr)

	// Parsing a string to a floating-point number
	floatStr := "3.14159"
	float, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("String %s parsed to float %f\n", floatStr, float)

	// Formatting an integer to a binary string
	binaryStr := strconv.FormatInt(int64(numInt), 2)
	fmt.Printf("Integer %d formatted to binary string %s\n", numInt, binaryStr)
}

