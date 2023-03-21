# Slices

A slice is a dynamically sized, flexible view of an underlying array. It is a reference to a ***contiguous segment of an array***, and can be created using the built-in `make` function. Slices are ***similar to arrays but provide more flexibility*** in terms of size, and they are passed by reference rather than by value.

A slice has three major components:

1. **Pointer**: A slice is implemented as a data structure that contains a pointer to the underlying array that holds the elements of the slice. The pointer points to the first element of the slice.
2. **Length**: The length of a slice is the number of elements in the slice. It represents the number of elements that are accessible through the slice. the `len()`
 function is used to get the length of the slice.
3. **Capacity**: The capacity of a slice is the maximum number of elements that the underlying array can hold without reallocating memory. It represents the total number of elements that can be stored in the slice. `cap()` function is used to get the capacity of the slice.

    ```go
    s := make([]int, 5, 10)
    fmt.Println(len(s)) // Output: 5
    fmt.Println(cap(s)) // Output: 10
    ```

> 📌 Note that the capacity of a slice can be greater than its length, but it can never be smaller. When the length of a slice is increased beyond its capacity, a new underlying array is allocated with a larger capacity, and the elements from the old array are copied to the new array.

A slice does not store any data, it just describes a section of an underlying array.

Here the length of the slice is 4, which means that there are 4 elements accessible through the slice. The capacity of the slice is 5, which means that the underlying array can hold up to 5 elements without reallocating memory.

```
	   +-----+-----+-----+-----+-----+-----+-----+----+
      |  0  |  1  |  2  |  3  |  4  |  5  |  6  |  7  |
      +-----+-----+-----+-----+-----+-----+-----+-----+
         ^                          ^                 ^
         |                          |                 |
      Pointer                    Length(4)        Capacity(8)

In this diagram, the pointer points to the first element of the underlying array. The length of the slice is 4, which means that there are 4 elements accessible through the slice. The capacity of the slice is 8, which means that the underlying array can hold up to 8 elements without reallocating memory.
```

### 1. Creating a Slice
    
```go
// Create an empty slice with a length of 0 and a capacity of 3
s1 := make([]int, 0, 3)

// Create a slice with initial values
s2 := []int{1, 2, 3}
```
    
### 2. Accessing elements in a slice:
    
```go
// Get the value at index 1
value := s2[1]
```
    
### 3. Modifying elements in a slice:
    
```go
// Set the value at index 1 to 4
s2[1] = 4
```

Changing the elements of a slice modifies the corresponding elements of its underlying array.</br>
Other slices that share the same underlying array will see those changes.

```go
    names := [4]string{"John","Paul","George","Ringo"}
    fmt.Println(names) //[John Paul George Ringo]

    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)  //[John Paul] [Paul George]

    b[0] = "XXX"
    fmt.Println(a, b)  //[John XXX] [XXX George]
    fmt.Println(names) //[John XXX George Ringo]
```
    
### 4. Slicing a slice:
    
```go
//newSlice := slice[start:end]
numbers := []int{1, 2, 3, 4, 5}
newSlice := numbers[1:3] // newSlice contains [2, 3]
newSlice1 := numbers[:3] // newSlice1 contains [1, 2, 3]
newSlice2 := numbers[3:] // newSlice2 contains [4, 5]
```
    

### 5. Appending to a slice:
    
```go
// Append a value to the end of a slice
s1 = append(s1, 1)

// Append multiple values to the end of a slice
s1 = append(s1, 2, 3, 4)

slice1 := []int{1, 2, 3}
slice2 := append(slice1, 4, 5)
    ```
    
### 6. Copying a slice:
    
```go
// Create a new slice with the same length and capacity as the original slice
s5 := make([]int, len(s2), cap(s2))

// Copy the values from s2 into s5
copy(s5, s2)
```
    
### 7. Creating a slice from an array
**`array[start_index:end_index]`**</br>
where `array` is the original array, `start` is the index of the first element in the slice, and `end` is the index of the element just after the last element in the slice.

```go
// slice := array[start:end]
numbers := [5]int{1, 2, 3, 4, 5}
slice := numbers[1:3] // slice contains [2, 3]
```

The `end` index is exclusive, so in this example, the slice contains elements at index 1 and 2, but not 3.