package main

import "fmt"

func main() {
	fmt.Println("Escape Sequences in Go:\n")

	// New Line
	fmt.Println("This is on the first line.\nThis is on the second line.")

	// Tab
	fmt.Println("This is some text.\tThis is indented by a tab.")

	// Backspace
	fmt.Println("This is some text.\b\b\b\b\b\bThis is overwritten by backspaces.")

	// Carriage Return
	fmt.Println("This is some text.\rThis overwrites the beginning of the line.")

	// Double Quote
	fmt.Println("\"This is in double quotes.\"")

	// Single Quote
	fmt.Println("'This is in single quotes.'")

	// Backslash
	fmt.Println("This uses a backslash: \\")

	// Null Character
	fmt.Println("This is before a null character.\x00This is after a null character.")
}
