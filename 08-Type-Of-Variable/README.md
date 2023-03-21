# Finding the type of variable

`fmt.Printf` function along with the `%T` format specifier to print the type of a variable.

2nd method is to use the `reflect` package to determine the type of a variable at runtime.

```go
package main

import "fmt"

func main() {
    var x int
    var y float64
    var z string

		// fmt.Printf function along with the %T format specifier
    fmt.Printf("x is of type %T\n", x)
    fmt.Printf("y is of type %T\n", y)
    fmt.Printf("z is of type %T\n", z)

		// using reflect package
    fmt.Println(reflect.TypeOf(x))
    fmt.Println(reflect.TypeOf(y))
    fmt.Println(reflect.TypeOf(z))

}
```

**Output:**

```go
x is of type int
y is of type float64
z is of type string

int
float64
string
```

The `%T` format specifier prints the type of the corresponding argument. This can be useful when you need to debug or check the type of a variable in your program.

The `reflect.TypeOf` function returns the type of its argument as a `reflect.Type` value. This can be useful when you need to determine the type of a variable at runtime, for example when working with interfaces or when you need to dynamically create values of different types.