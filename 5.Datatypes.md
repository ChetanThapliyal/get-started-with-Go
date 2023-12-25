# Data Types

## What is Data Type and Why are they needed?

* Data type is a classification of data that defines the type of data it is.
* Data types are used to define the operations that can be performed on the data, the meaning of the data, and the way the data is stored.
* Common data types include integers, strings, booleans, floats, arrays and slices.
* Each data type can have different sizes, range of values, and operations that can be performed on it.

> ### 📌 What is Memory Allocation?
> - Memory allocation refers to the process of reserving a portion of memory for storing data. 
> - When a variable is declared in a program, the system allocates a certain amount of memory to store the value of that variable. 
> - The size of the memory block depends on the data type of the variable.

For example, if you declare an `int` variable, the system will allocate 4 bytes of memory to store the integer value. Similarly, if you declare a `float32` variable, the system will allocate 4 bytes of memory to store the floating-point value.

> It's important to note that the ***amount of memory required by different data types can vary between different computer architectures and operating systems***. However, Go provides a standardized way of allocating memory to different data types on all platforms.

## Data Types in GO
Go supports a range of data types, including numeric, boolean and string types, as well as composite types such as arrays, slices, maps and structures.

### Numeric Types

Go’s numeric types include `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, and `complex64`, `complex128`.

`uint` = unsigned integer (only positive no.’s inc. 0)

Example:

```go
// Declare variables with int type
var age int = 22

// Declare variables with float type
var pi float32 = 3.14

// Declare variables with complex type
var c complex64 = 2.3 + 2.4i
```

### Boolean Types

In Go, the boolean type is `bool`. It can only contain two values: `true` or `false`

Example:

```go
// Declare a boolean variable
var isValid bool = false
```

### String Types

Go supports strings of type `string`.

Example:

```go
// Declare a string variable
var name string = "John"
```

### Array Types

Go supports arrays of type `[size]T`, where `size` is the number of elements in the array, and `T` is the type of each element.

Example:

```go
// Declare an array of type int
var numbers [3]int = [3]int {1,2,3}
cgpa := [3]int {7,8,9}

//use the ellipsis notation to let Go determine the size of the array automatically based on the number of values provided:
marks := [...]{93,86,78.99.87}
```

### Slice Types

Go also supports slices of type `[]T`, where `T` is the type of each element.

Example:

```go
// Declare a slice of type int
var numbers []int = []int {1,2,3}
```

### Map Types

Go also supports maps of type `map[K]V`, where `K` is the type of the key, and `V` is the type of the value.

Example:

```go
// Declare a map of type int and string
var numbersMap map[int]string = make(map[int]string)
numbersMap[1] = "John"
numbersMap[2] = "Mary"
```

### Struct Types

Go also supports structures of type `struct`. Structures are used to group related data together.

Example:

```go
// Declare a struct of type Person
type Person struct {
	name string
	age  int
}

// Declare a variable of type Person
var john Person = {name: "John", age: 22}
```