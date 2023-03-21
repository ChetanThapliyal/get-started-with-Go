package main

import "fmt"

func main() {

	//INITIALIZING

	var a [5]int // creates an integer array with length 5

	// To access an element of an array, you need to specify its index value
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50

	fmt.Println(a[0]) // Output: 10

	// You can also initialize an array when you declare it:
	b := [5]int{60, 70, 80, 90, 100} // var b [5]int = [5]int{60, 70, 80, 90, 100}
	fmt.Println(b[3])                // Output: 90

	// Or, you can use the ellipsis notation to let Go determine the size of the array automatically based on the number of values provided:
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c[4]) // Output: 5

	//Changing value of elements in array:
	fmt.Println(a) //Output:[10 20 30 40 50]
	a[3] = 60
	fmt.Println(a) //Output:[10 20 30 60 50]

	//Multi Dimentional array
	var arr [2][3]int //2D array with 2 rows and 3 columns
	arr[0][0] = 1
	arr[0][1] = 2
	arr[0][2] = 3

	arr[1][0] = 4
	arr[1][1] = 5
	arr[1][2] = 6

	fmt.Println(arr)

	//shorthand method of initializing
	arrr := [2][3]int{{7, 8, 9}, {10, 11, 12}}
	fmt.Println(arrr) //Output: [[7 8 9] [10 11 12]]

	//Array with 2 planes, 3 rows, and 4 columns
	array := [2][3][4]int{{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, {{13, 14, 15, 16}, {17, 18, 19, 20}, {21, 22, 23, 24}}}

	fmt.Println(array) //Output:[[[1 2 3 4] [5 6 7 8] [9 10 11 12]] [[13 14 15 16] [17 18 19 20] [21 22 23 24]]]

}

