package main

import "fmt"

func main() {

	//declaration and initialization of slice
	// Create an empty slice with a length of 0 and a capacity of 3
	s1 := make([]int, 0, 3)
	fmt.Println(s1)
	fmt.Println(cap(s1))
	fmt.Println(len(s1))

	// Create a slice with initial values
	s2 := []int{1, 2, 3}
	fmt.Println(s2)

	// Get the value at index 1
	value := s2[1]
	fmt.Println(value)

	//Creating a slice from an array
	// slice := array[start:end]
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]
	fmt.Println(slice)

	// Modifying values
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names) //[John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) //[John Paul] [Paul George]

	b[0] = "XXX"
	fmt.Println(a, b)  //[John XXX] [XXX George]
	fmt.Println(names) //[John XXX George Ringo]

	//slicing a slice
	//newSlice := slice[start:end]
	numbers := []int{1, 2, 3, 4, 5}
	newSlice := numbers[1:3]
	newSlice1 := numbers[:3]
	newSlice2 := numbers[3:]
	fmt.Println(newSlice)
	fmt.Println(newSlice1)
	fmt.Println(newSlice2)

	//Appending a slice
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1)
	fmt.Println(slice2)

	//Copying a slice
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, len(slice3))
	copy(slice4, slice3)
	fmt.Println(slice3)
	fmt.Println(slice4)

}

