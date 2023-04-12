package main

import "fmt"

func main() {
	// declare a variable
	x := 1

    //Initializing and declaring pointer that points to the memory address of x
	var ptr *int = &x
    
	/* or compiler can determine datatype internally
	   var ptr =&x
	   or using shorthand declaration
	   ptr := &x 
	*/

	// print the value of x
	fmt.Println("x =", x)

	// print the memory address of x and datatype
	fmt.Println("Memory address of x =", &x)
	fmt.Printf("Datatype of &x = %T \n", &x)

	// print the value of the pointer variable (which is the memory address of x)
	fmt.Println("Pointer variable value =", ptr)

	// dereferencing the pointer variable to get the value of x and datatype
	fmt.Println("Dereferenced pointer variable value =", *ptr)
	fmt.Printf("Datatype of  dereference pointer *(&x) = %T \n", *(&x))

	// update the value of x through the pointer variable
	*ptr = 20
	fmt.Println("Updated value of x =", x)
}