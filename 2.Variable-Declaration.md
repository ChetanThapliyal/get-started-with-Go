# Variables

- A named storage location in a computer program that holds a value of a specific data type.
- The value of a variable can be changed during the execution of a program, and the variable can be used to store data that will be used later in the program.
- A variable has a name, a data type, and a value.
- The name of the variable is used to refer to it in the program, and the data type specifies the type of value that the variable can hold.
- The value of the variable can be assigned at the time of its declaration or later during the execution of the program.
- Variables are an important concept in programming because they allow programs to store and manipulate data dynamically. They also provide a way to reuse data and avoid hardcoding values in the program, which can make it more flexible and easier to modify.

## Different ways of Declaring Variables in Go
In Go, there are several ways to declare variables, including the following:

1. `var` keyword: This is the most common way to declare a variable in Go. It requires that the data type of the variable be specified:
    
    ```go
    // var variable_name data_type = value 
    var age int = 26
    ```
    
2. `:=` shorthand: This is a shorthand way of declaring and initializing a variable in Go. The data type of the variable is inferred based on the value assigned:
    
    ```go
    name := "John Doe"
    ```
    
3. `var` block: This allows you to declare multiple variables in a single block:
    
    ```go
    var (
      name string = "John Doe"
      age int = 26
    )
    ```
    

4. `const` keyword: This is used to declare constant values in Go, which cannot be changed once they are assigned:
    
    ```go
    const Pi float64 = 3.14
    ```
    
    It's also possible to declare multiple constants in a single block:
    
    ```go
    const (
      E float64 = 2.718
      Pi float64 = 3.14
    )
    ```
    

> In Go, it's important to choose the right way to declare variables based on your specific needs and goals. 

For example, if you want to declare a constant value, you should use the `const` keyword. 

On the other hand, if you want to declare a variable that will change its value, you should use the `var` keyword or the `:=` shorthand.