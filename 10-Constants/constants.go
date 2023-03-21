package main

func main() {

	/*----------------
	Typed Constants
	-----------------*/
	// Define integer constant
	const a int = 10

	// Define float constant
	const b float64 = 3.14

	// Define string constant
	const c string = "Hello, world!"

	// Define boolean constant
	const d bool = true

	/*----------------
	Untyped Constants
	-----------------*/

	const e = 10
	const f = "hello"

	/*-------------------------------
	Difference b/w Typed and Untyped
	--------------------------------*/

	const typedConst int = 10
	const untypedConst = 10

	var g int = typedConst / 3   // valid arithmetic operation
	var h int = untypedConst / 3 // valid arithmetic operation

	var i float64 = typedConst / 3   // invalid arithmetic operation, different types
	var j float64 = untypedConst / 3 // valid arithmetic operation, inferred as float64

	var k bool = typedConst == 10   // valid comparison, same types
	var l bool = untypedConst == 10 // valid comparison, inferred as int

}

