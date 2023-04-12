# Introduction

Pointers are variables that store the memory address of another variable. They are represented using the `*` symbol. 

Pointers are commonly used in Go to achieve pass-by-reference behavior, which allows functions to modify the original values of variables passed to them.

```go
func main() {
    // declare a variable
    x := 1

    // declare a pointer variable that points to the memory address of x
    var ptr *int = &x

    // print the value of x
    fmt.Println("x =", x)

    // print the memory address of x
    fmt.Println("Memory address of x =", &x)
}
```

## Address and Dereference Operator
The memory address is represented by a hexadecimal number, which is a base 16 number system. The address of a variable can be obtained by using the **`&`** operator followed by the variable name.

For example, if we have a variable `x` of type `int`, we can get its memory address by using `&x`. The `&` operator is known as the ***address operator***.

Here's an example:

```go
x := 10
fmt.Println("Value of x:", x)  // prints: Value of x: 10
fmt.Println("Address of x:", &x) // prints: Address of x: 0xc0000180a0
fmt.Printf("Type of x: %T", &x)
```

In the above example, `x` holds the value `10` and `&x` holds the memory address of `x`, which is `0xc0000180a0`.

To access the value at the memory address held by a pointer variable, we use the `*` operator. This is known as the ***dereference operator***.

Here's an example:

```go
x := 10
ptr := &x
fmt.Println("Value of x:", x)   // prints: Value of x: 10
fmt.Println("Value of ptr:", *ptr) // prints: Value of ptr: 10
fmt.Printf("Type of x: %T", &x)
```

In the above example, `ptr` holds the memory address of `x`, which we obtained using the `&` operator. 

To access the value at the memory address held by `ptr`, we use the `*` operator. This gives us the value of `x`, which is `10`.

```go
package main

import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))
}
```

> ⭐ Note that when using pointers in Go, it's important to handle **nil pointers** properly to avoid runtime errors.

## Passing by Value in Functions

In Go, all function arguments are passed by value, meaning a copy of the argument is passed to the function. However, when the argument is a pointer, a copy of the memory address is passed, allowing the function to modify the original variable.

> - Function is called by directly passing the value of the variable as an argument. The parameter is copied into another location of your memory.
> - So, when accessing or modifying the variable within your function, only the copy is accessed or modified, and the original value is never modified.

Here's an example of passing by value:

```go
func main() {
    x := 10
    increment(x)
    fmt.Println(x) // prints: 10
}

func increment(a int) {
    a++
}
```

In the above example, the `increment` function takes an argument `a` of type `int`. When we call `increment(x)`, a copy of `x` is passed to the function. The `increment` function increments the value of `a`, but this has no effect on the original variable `x` because it's a copy.

## Passing by Reference in Functions

Here's an example of passing by reference using a pointer:

```go
func main() {
    x := 10
    increment(&x)
    fmt.Println(x) // prints: 11
}

func increment(a *int) {
    (*a)++
}
```

In the above example, we pass the memory address of `x` to the `increment` function using the `&` operator. The `increment` function takes a pointer argument `a` of type `*int`. To access the value at the memory address held by `a`, we use the `*` operator, which dereferences the pointer.

When we call `increment(&x)`, the `increment` function modifies the value at the memory address held by `a`, which is the same memory address held by `x`. Therefore, the value of `x` is also incremented to `11`.

> Note that using pointers in Go can be useful for passing large data structures or modifying function arguments in place. However, it also requires careful handling to avoid errors such as nil pointers or memory leaks.