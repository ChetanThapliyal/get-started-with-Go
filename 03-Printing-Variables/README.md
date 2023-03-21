# Methods of Printing variables

In Go, there are several methods of printing variables, including the following:

1. `fmt.Println`: This is the most commonly used method of printing variables in Go. It prints the value of the variable to the console and ***adds a newline character at the end***:
    
    ```go
    var name string = "Harry Potter"
    var house string = "Gryffindor"
    fmt.**Println**(name)
    fmt.**Println**(house)
    
    // OUTPUT: 
    // Harry Potter
    // Gryffindor 
    ```
    
2. `fmt.Printf`: This is used to ***print a formatted string to the console***, which can include variables. It supports various formatting options, such as 
    - `%v`: Prints the value in a default format.
    - `%d`: Prints an integer value in decimal format.
    - `%f`: Prints a floating-point value in decimal format.
    - `%s`: Prints a string value.
    - `%t`: Prints a boolean value.
    - `%c`: Prints a character value.
    
    >    📌 In Go, format specifiers are special characters that are used to specify the format of a value when printing it to the console or to a string. Format specifiers are used with the `fmt.Printf` function or with other functions that support formatted printing.
    
    ```go
    name := "John Doe"
    fmt.Printf("My name is %s\n", name)
    
    // >>> output
    // My name is John Doe
    ```
    
3. `fmt.Sprintf`: This is used to format a string and return it as a string value. This can be useful if you want to store the formatted string in a variable or use it in other parts of your program:
    
    ```go
    age := 26
    message := fmt.Sprintf("My age is %d", age)
    fmt.Println(message)
    
    // >>> output
    // My age is 26
    ```
    
    In addition to these methods, there are also other packages in Go that can be used to print variables, such as the `log` package and the `strconv` package for converting numbers to strings. The choice of method often depends on the specific needs and goals of the program.