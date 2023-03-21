# Type Casting : Converting type of Variables

Type conversion is the process of converting a value of one type to another type.

## Explicit Type Casting
This type of type conversion is done explicitly by the programmer using the syntax T(expression), where T is the type to which you want to convert the expression.

By using the syntax `T(value)` to convert the value of type `value` to type `T`. 

```go
package main

import "fmt"

func main() {
    // Convert an int to a float64
    x := 10 //var x int =10
    y := float64(x) //var y float64 = float64(x)
    fmt.Println(y)

    // Convert a float64 to an int
    z := 10.5
    w := int(z)
    fmt.Println(w)

    // Convert a string to a byte slice
    s := "hello"
    b := []byte(s)
    fmt.Println(b)

    // Convert a byte slice to a string
    b2 := []byte{'h', 'e', 'l', 'l', 'o'}
    s2 := string(b2)
    fmt.Println(s2)
}
```

**Output:**

```go
10
10
[104 101 108 108 111]
hello
```


> 📌 Note that **not all types can be converted to all other types**. 
For example, you cannot convert a string to an int directly, or vice versa.</br>
> In addition, type conversion can result in loss of precision or truncation of values in certain cases, so you should be careful when performing type conversion.

## Implicit Type Conversion 
This type of type conversion is automatically done by the Go compiler. For example, if you assign an int value to a float64 variable, the compiler will automatically convert the int to float64.

```go
var i int = 42
var f float64 = i   // The compiler will automatically convert i to float64
fmt.Println(f)      // Output: 42.0
```

## `strconv` package

`strconv` package provides functions for converting strings to numeric types and vice versa. It is useful when dealing with input/output or configuration files where values are typically read in as strings.

Some of the important functions in the `strconv` package include:

1. **`Atoi`** and **`ParseInt`**:</br> 
    These functions convert a _string to an integer_.</br>
    `Atoi` is a shorthand for `ParseInt(s, 10, 0)`, which converts a string `s` to a base-10 integer with a default bit size of 0.</br>
    The `ParseInt` function takes three arguments: the string to convert, the base (usually 10), and the bit size (32, 64, or 0 for the size of the platform's native int type).</br>

2. **`Itoa`** and **`FormatInt`**:</br>
These functions convert an _integer to a string_.</br>
`Itoa` is a shorthand for `FormatInt(int64(i), 10)`, which converts an integer `i` to a base-10 string.</br>
The `FormatInt` function takes the integer to convert, the base (usually 10), and the bit size.</br>

3. **`ParseFloat`** and **`FormatFloat`**:</br>
These functions convert a _string to a float64_ and vice versa.</br>
`ParseFloat` takes three arguments: the string to convert, the bit size (32 or 64), and the precision (the number of bits to use for the mantissa).</br>
`FormatFloat` takes the float to convert, the format specifier, and the precision.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Converting a string to an integer
	numStr := "42"
	num, _ := strconv.**Atoi**(numStr)		
	fmt.Printf("String %s converted to integer %d\n", numStr, num)

	// Converting an integer to a string
	numInt := 12345
	numStr = strconv.**Itoa**(numInt)
	fmt.Printf("Integer %d converted to string %s\n", numInt, numStr)

	// Parsing a string to a floating-point number
	floatStr := "3.14159"
	float, _ := strconv.**ParseFloat**(floatStr, 64)
	fmt.Printf("String %s parsed to float %f\n", floatStr, float)

	// Formatting an integer to a binary string
	binaryStr := strconv.**FormatInt**(int64(numInt), 2)
	fmt.Printf("Integer %d formatted to binary string %s\n", numInt, binaryStr)
}
```

In the above example, `Atoi` function is used to convert a string representation of an integer to an actual integer.

`Itoa` function is used to convert an integer to its string representation. 

`ParseFloat` function is used to parse a string representation of a floating-point number to its actual value. 

`FormatInt` function is used to format an integer to its binary string representation.

## `reflect` Package 
The reflect package in Go provides a way to examine and manipulate the types of values at runtime. Using this package, you can convert a value to another type using the `Value.Convert()` method. However, this method is slower than the other methods and should be used sparingly.

```go
import "reflect"

var i int = 42
v := reflect.ValueOf(i)  // Creating a reflect.Value from i
f := v.Convert(reflect.TypeOf(float64(0))).Float()  // Converting v to a float64 value
fmt.Println(f)           // Output: 42.0
```
Note that in above example, we had to provide a target type `(float64(0))` to the `Convert()` method in order to convert the value `v` to the desired type.