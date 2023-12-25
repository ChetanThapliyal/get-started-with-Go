# Arrays

An array is a collection of elements of the **same data type** stored in contiguous memory locations. It is a **fixed-size collection** of elements, and each element can be accessed using an index value. In Go, arrays **can be of any data type** and have a fixed size, which cannot be changed at runtime.

Arrays can be passed to functions as arguments, but note that passing an array to a function creates a copy of the array, which can be memory-intensive for large arrays. Here's an example of passing an array to a function:

```go
func sum(numbers [5]int) int {
    result := 0
    for _, number := range numbers {
        result += number
    }
    return result
}

func main() {
    a := [5]int{10, 20, 30, 40, 50}
    fmt.Println(sum(a))
}
```

Arrays in Go are quite basic and limited compared to other data structures like slices, which can grow or shrink as needed. However, they can be useful in certain cases where a fixed-size collection of elements is required.

## Array Indexes 

array indexes start from 0 and go up to length-1. For example, if an array has a length of 5, its indexes will be 0, 1, 2, 3, and 4. Here's an example of how you can access array elements using their indexes:

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr[0]) // Output: 1
    fmt.Println(arr[2]) // Output: 3
    fmt.Println(arr[4]) // Output: 5
}
```

Changing value of elements in array:

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr) // Output: 1,2,3,4,5
    arr[1] = 7
		fmt.Println(arr) // Output: 1,**7**,3,4,5
}
```

```go
var a [5]int // creates an integer array with length 5

// To access an element of an array, you need to specify its index value
a[0] = 10
a[1] = 20
a[2] = 30
a[3] = 40
a[4] = 50

// You can also initialize an array when you declare it:
b := [5]int{60, 70, 80, 90, 100} // var b [5]int = [5]int{60, 70, 80, 90, 100}

// Or, you can use the ellipsis notation to let Go determine the size of the array automatically based on the number of values provided:
c := [...]int{1, 2, 3, 4, 5}
```

Iterating over an array:

```go
//iterating over an array:
a := [...]int{1, 2, 3, 4, 5}
for i := 0; i < len(a); i++ {     //len(): length function 
    fmt.Println(a[i])
}

//can also use a range loop to iterate over an array:
for index, value := range a {     //range: iterate over all elements
    fmt.Println(index, value)
}
```

```go
arr[0][1] = 42
```

## Multi Dimentional array

A multidimensional array is an array in which each element is itself an array. The number of dimensions of a multidimensional array is defined by the number of nested arrays.

To create a multidimensional array in Go, we can define an array whose elements are also arrays. For example, the following code creates a 2D array with 2 rows and 3 columns:

```go
var arr [2][3]int
```

We can access the elements of a multidimensional array using the index of the element in each dimension. For example, to access the element in the first row and second column of the above array, we would use:

```go
arr[0][1] = 42

//OR
//arr := [2][3]int{{7,8,9},{10,11,12}}
```

An example of a 3D array with 2 planes, 3 rows, and 4 columns:

```go
var arr [2][3][4]int
```

We can access elements of a 3D array using three indices, one for each dimension. For example, to set the element in the second plane, first row, and third column of the above array to 10, we would use:

```go
arr[1][0][2] = 10
```

Example that creates a 2D array and sets some of its elements:

```go
package main

import "fmt"

func main() {
    var arr [2][3]int
    
    arr[0][0] = 1
    arr[0][1] = 2
    arr[0][2] = 3
    
    arr[1][0] = 4
    arr[1][1] = 5
    arr[1][2] = 6
    
// OR

		arrr := [2][3]int{{7,8,9},{10,11,12}}
    
		fmt.Println(arr)
		fmt.Println(arrr)
}
```

Output:

```go
[[1 2 3] [4 5 6]]
[[7,8,9] [10,11,12]]
```