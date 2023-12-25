# Control Flow

Control flow refers to the order in which the statements in a program are executed. In Go, there are several control flow statements that allow you to control the flow of execution in a program. These include:

## If/Else statements

An if statement is used to execute a block of code if a certain condition is true. 

An else statement can be added to execute a different block of code if the condition is false.

```go
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}
```

## Switch statements

A switch statement is used to compare a value against multiple possible values and execute different blocks of code depending on which value matches.

```go
package main

import "fmt"

func main() {
    day := "Tuesday"

    switch day {
    case "Monday":
        fmt.Println("Today is Monday")
    case "Tuesday":
        fmt.Println("Today is Tuesday")
    case "Wednesday":
        fmt.Println("Today is Wednesday")
    default:
        fmt.Println("Today is another day")
    }
}
```

## Break and continue statements

The break statement is used to exit a loop prematurely, while the continue statement is used to skip the current iteration and move on to the next one.

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            break
        }
        if i % 2 == 0 {
            continue
        }
        fmt.Println(i)
    }
}
```

## While loops

`while`
 loop is a control structure that repeatedly executes a block of code as long as a specified condition is true. The syntax for a `while` loop in Go is as follows:

```go
//Syntax
for condition {
    // code to be executed while the condition is true
}
```

The `condition` is an expression that is evaluated at the beginning of each iteration of the loop. If the condition is true, the code inside the loop is executed. If the condition is false, the loop terminates and the program continues with the code that follows the loop.

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

In this example, the `while` loop is used to print the numbers from 0 to 4. The condition `i < 5` is evaluated at the beginning of each iteration of the loop. As long as `i` is less than 5, the code inside the loop is executed, which prints the value of `i` and increments it by 1. When `i` reaches 5, the condition is no longer true and the loop terminates.

## For loops

A for loop is used to repeat a block of code a fixed number of times.

```go
//SYNTAX
for initialization; condition; post {
    // Code to be executed
}
```

- The `initialization` statement is executed before the first iteration of the loop. It is typically used to initialize a variable that will be used as a loop counter, or to perform some other initialization task.
- The `condition`  statement is evaluated at the beginning of  each iteration of the loop. If the condition is true, the loop body is  executed; otherwise, the loop exits.
- The `post` statement is executed at the end of each iteration of the loop, after the loop body has been executed. It is typically used to update the loop counter or perform some other action.

Here's an example of a simple for loop that prints the numbers 0 to 9:

```go
for i := 0; i < 10; i++ {
       fmt.Println(i) }
```

In this example, the `initialization` statement initializes the variable `i` to 0. The `condition` statement checks if `i` is less than 10. If it is, the loop body is executed, which prints the value of `i` to the console using the `fmt.Println` function. The `post` statement increments the value of `i` by 1.

The output of this code would be: 0, 1, 2, 3, 4, 5, 6, 7, 8, 9

### `range` keyword
**`Range`** Keyword : the `range` keyword is used to iterate over the elements of an array, slice, map, or channel. The basic syntax of the `range` keyword in Go is as follows: 

```go
for index, value := range collection {
// Code to be executed
}
```

Let's say we have a slice of strings:

```go
fruits := []string{"apple", "banana", "cherry"}
for index, value := range fruits {
    fmt.Println(index, value)
}
```

We can use a `range` loop to iterate over the elements of the `fruits` slice and print them to the console:

```go
for index, value := range fruits {
    fmt.Println(index, value)
}
```

In this code, the loop will iterate over the elements of the `fruits` slice, and at each iteration, the loop body will print the current index and value to the console using the `fmt.Println` function.

The output of this code would be:

```go
0 apple
1 banana
2 cherry
```

## Fallthrough, Break and Continue Statements

`fallthrough`, `break` and `continue` statements are used in control flow constructs to modify the control flow of a program.

### `fallthrough`

- `fallthrough` statement: It is used in a `switch` statement to transfer control to the next `case` block. When `fallthrough` is encountered in a `case` block, control passes to the next `case` block, regardless of whether the condition for the `case` block is met or not. Here is an example:

```go
package main

import "fmt"

func main() {
    i := 2
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
        fallthrough
    case 3:
        fmt.Println("three")
    }
}
```

In this example, when the `i` variable is equal to `2`, the program prints "two" and then falls through to the next case block and prints"three".

```go
// OUTPUT
two
three
```

### `break`

`break` statement: It is used in a loop or `switch` statement to break out of the loop or switch statement. When `break` is encountered in a loop or `switch` statement, the loop or switch statement is exited immediately. Here is an example:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            break
        }
        fmt.Println(i)
    }
}
```

In this example, the program prints the value of `i` until `i` is equal to `5`. When `i` is equal to `5`, the `break` statement is encountered and the loop is exited immediately.

```go
//OUTPUT
1
2
3
4
```

### `continue`

`continue` statement: It is used in a loop to skip the rest of the loop body for the current iteration and move on to the next iteration of the loop. When `continue` is encountered in a loop, the loop immediately starts the next iteration. Here is an example:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            continue
        }
        fmt.Println(i)
    }
}
```

In this example, the program prints the value of `i` except when `i` is equal to `5`. When `i` is equal to `5`, the `continue` statement is encountered and the loop immediately starts the next iteration.

```go
//OUTPUT
1
2
3
4
6
7
8
9
10
```