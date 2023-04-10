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