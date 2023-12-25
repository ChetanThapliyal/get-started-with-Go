# Functions

Functions are generally the block of codes or statements in a program that gives the user the ability to reuse the same code which ultimately saves the excessive use of memory, acts as a time saver and more importantly, provides better readability of the code. So basically, a function is a collection of statements that perform some specific task and return the result to the caller. A function can also perform some specific task without returning anything.

### Function Declaration

Function declaration means a way to construct a function.</br>
**Syntax:** 

```go
func function_name(Parameter-list)(Return_type){
    // function body.....
}
```

The declaration of the function contains:

- **func:** It is a keyword in Go language, which is used to create a function.
- **function_name:** It is the name of the function.
- **Parameter-list:** It contains the name and the type of the function parameters.
- **Return_type:** It is optional and it contain the types of the values that function
returns. If you are using return_type in your function, then it is necessary to use a return statement in your function.

    ```go
    func add(x int, y int) int {
        return x + y
    }

    add(2,3)
    ```

### Naming Conventions of function

- must begin with a letter.
- can have any number of additional letters and symbols.
- cannot contain spaces.
- case-sensitive.


### Calling a Function

```go
<function_name>(<argument(s)>)
```

```go
package main
import "fmt"
 
// area() is used to find the
// area of the rectangle
// area() function two parameters,
// i.e, length and width
func area(length, width int)int{
     
    Ar := length* width
    return Ar
}
 
// Main function
func main() {
   
   // Display the area of the rectangle
   // with method calling
   fmt.Printf("Area of rectangle is : %d", area(12, 10))
}
```

### Parameters vs Arguments

- A "**parameter**" is a variable or a value that is part of the function or method's definition or signature, specifying what type and how many inputs it expects. Parameters are often listed in the function or method declaration, and they serve as placeholders for the actual input values that will be provided when the function or method is called.
- An "**argument**," on the other hand, is the actual value or expression that is passed to a function or method when it is called. Arguments are the inputs that are provided to the function or method at runtime.

To illustrate the difference between parameters and arguments, consider the following function definition:

```go
func sum(a int, b int) int {
    return a + b
}
```

In this function definition, `a` and `b` are the parameters, and they define what inputs the function expects. When calling this function and passing in actual values, those values are the arguments:

In this example, `3` and `4` are the arguments, and they are passed to the function to be used as inputs.

```go
total := sum(3, 4)
```

## Go language supports two ways to pass arguments to your function:
---
**Call by value:** In this parameter passing method, values of actual parameters are copied to function’s formal parameters and the two types of parameters are stored in different memory locations. So any changes made inside functions are not reflected in actual parameters of the caller. **Example:**

```go
// Go program to illustrate the
// concept of the call by value
package main
import "fmt"

// function which swap values
func swap(a, b int)int{

    var o int
    o= a
    a=b
    b=o
    
return o
}

// Main function
func main() {
var p int = 10
var q int = 20
fmt.Printf("p = %d and q = %d", p, q)

// call by values
swap(p, q)
fmt.Printf("\np = %d and q = %d",p, q)
}
```

**Output:**

```go
p = 10 and q = 20
p = 10 and q = 20
```

**Call by reference:** Both the actual and formal parameters refer to the same locations, so any changes made inside the function are actually reflected in actual parameters of the caller. **Example:**

```go
// Go program to illustrate the
// concept of the call by reference
package main
import "fmt"

// function which swap values
func swap(a, b *int)int{
	var o int
	o = *a
	*a = *b
	*b = o
	
return o
}

// Main function
func main() {

var p int = 10
var q int = 20
fmt.Printf("p = %d and q = %d", p, q)

// call by reference
swap(&p, &q)
fmt.Printf("\np = %d and q = %d",p, q)
}
```

**Output:**

```go
p = 10 and q = 20
p = 20 and q = 10
```

## Return Type Function : Multiple Named, Variadic

In Golang, a function can have a return type that specifies the type of value that the function returns. The return type is declared after the list of parameters, if any, and before the opening curly brace of the function body.

**Returning Single and Multiple value**

```go
package main
import "fmt"

// Function with single return type
func addNumbers(x int, y int) int {
    return x + y
}

// Function with a multiple returns
func operations(x int, y int) int {
	sum := x+y
    diff := x-y
	return sum, diff
}

func main() {
    result := addNumbers(5, 7)
    fmt.Println(result) // Output: 12

	sum, difference := operations(30,10)
	fmt.Println(sum, difference) // Output: 40, 20
}
```

**Variadic Functions**

- Function that accepts variable number of arguments. This allows you to write functions that can work with different numbers of arguments, without needing to write multiple versions of the function for each possible number of arguments.
- It is possible to pass a varying number of arguments of the same type as referenced in the function signature.
- To declare a variadic function, the type of the final parameter is preceded by an ellipsis "`. . .`"

    ```go
    //Syntax
    func function_name(param1 type, param2 type, Param3 ...)(Return_type){
        // function body.....
    }
    ```

    ```go
    package main

    import "fmt"

    // variadic
    func sumNum(num ...int) int {

        sum := 0
        for _, value := range num {
            sum += value
        }
        return sum
    }

    func main() {

        fmt.Println(sumNum())                    // Output: 0
        fmt.Println(sumNum(10))                  // Output: 10
        fmt.Println(sumNum(10, 20))              // Output: 30
        fmt.Println(sumNum(10, 20, 30, 40, 50))  // Output: 150
    }
    ```

**Named Return values**

- When a function returns multiple values, you can name the return values to make the code more readable and easier to understand.
- To name the return values in a function, you can simply declare the variable names after the function signature, separated by a comma.

    ```go
    package main
    import "fmt"

    func rectProps(length, width float64) (area, perimeter float64) {
        area = length * width
        perimeter = 2 * (length + width)
        return // returns both named values
    }

    func main() {
        area, perimeter := rectProps(5.0, 6.0)
        fmt.Println(area, perimeter) // Output: 40, 20
    }
    ```

## Recursive Function

- A recursive function is a function that calls itself during its execution until it reaches a base case where it stops recursing
- It is used to solve a problem where the solution is dependent on the smaller instance of the same problem.

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

`factorial()` is a recursive function that takes a non-negative integer `n` as an argument and returns the factorial of `n`.

The base case is when `n` is equal to 0, in which case the function returns 1. Otherwise, the function calls itself with the argument `n-1` and multiplies the result with `n` to get the factorial of `n`.

For example, `factorial(5)` would compute the factorial of 5 by calling `factorial(4)`, `factorial(3)`, `factorial(2)`, and `factorial(1)` until it reaches the base case of `factorial(0)`, which returns 1. 

The result is then multiplied by 2, then by 3, then by 4, and finally by 5 to get the final result of 120, which is the factorial of 5.

## Anonymous Function

- An anonymous function is a function without a name. Instead, it is declared inline using a function literal, which is a function definition that is not bound to an identifier. 

- Anonymous functions are useful when we need to define a function that is only used once or when we want to pass a function as an argument to another function.

- They can accept inputs and return outputs, just as standard functions do.

Here's an example of an anonymous function in Go:

```go
package main

import "fmt"

func main() {
	// Define and call an anonymous function
	func() {
		fmt.Println("Hello, anonymous function!")
	}()

	// Assign an anonymous function to a variable and call it
	add := func(a, b int) int {
		return a + b
	}

	sum := add(3, 4)
	fmt.Println(sum)
	fmt.Printf("%T \n", x)
}
```

In this example, we define and call an anonymous function on the first line of the `main` function. The function simply prints a message to the console.

On the following lines, we assign an anonymous function that takes two arguments and returns their sum to a variable named `add`. 

We then call the `add` function with the arguments `3` and `4` and store the result in a variable named `sum`. Finally, we print the value of `sum` to the console.

## High Order Functions

- A higher-order function is a function that takes one or more functions as arguments and/or returns a function as its result.
- Higher-order functions are a powerful tool for creating reusable and composable code, allowing developers to write more flexible and expressive programs.

**Why use High order Functions?**

- Composition
    - creating smaller functions that take care of certain piece of logic.
    - composing complex function by using different smaller functions.
- Reduces bugs
- Code gets easier to read and understand.

Here's an example of High Order function in Go:

```go
package main

import "fmt"

// greetFunction is a higher-order function that takes a string and returns another function.
func greetFunction(greeting string) func(string) string {
	// Return a new function that takes a string and returns a formatted greeting.
	return func(name string) string {
		return fmt.Sprintf("%s, %s!", greeting, name)
	}
}

func main() {
	// Create two functions that greet in different ways.
	englishGreeting := greetFunction("Hello")
	spanishGreeting := greetFunction("Hola")

	// Call the englishGreeting and spanishGreeting functions to get formatted greetings.
	fmt.Println(englishGreeting("Alice")) // Output: Hello, Alice!
	fmt.Println(spanishGreeting("Bob"))   // Output: Hola, Bob!
}
```

## Defer Statements
A `defer` statement is used to schedule a function call to be executed after the function that contains it completes. It is used to ensure that resources are released or functions are executed at the end of a function, regardless of how the function exits.

The `defer` statement is often used to pair an opening function call with a corresponding closing function call. For example, if a function opens a file, the `defer` statement can be used to ensure that the file is closed when the function returns, even if an error occurs.

Here's an example of how `defer` is used in Go:

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

In this example, the `fmt.Println("world")` statement is scheduled to be executed after the `fmt.Println("hello")` statement completes. When the program is run, it will output:

```go
hello
world
```

Another example could be a function that locks a mutex, and the corresponding `defer` statement that unlocks the mutex:

```go
func myFunc(mu *sync.Mutex) {
    mu.Lock()
    defer mu.Unlock()
    
    // rest of the function logic
}
```

In this example, the `mu.Lock()` statement locks the mutex, and the `defer mu.Unlock()` statement is scheduled to be executed after the function completes, ensuring that the mutex is always unlocked, even if an error occurs in the function.