# Variable Scope

Variable scope refers to the visibility or accessibility of a variable within a program. In Go, a variable's scope is determined by where it is declared in the program and can be categorized into two types:

1. **Local scope** : Variables declared within a function or a block have local scope and are accessible only within that function or block. Once the function or block is executed, the variable is destroyed and cannot be accessed anymore. For example:
    
    ```go
    func someFunction() {
        var x int // x has local scope
        x = 10
        fmt.Println(x) // prints 10
    }
    ```
    
2. **Global scope** : Variables declared outside of any function or block have global scope and can be accessed from any part of the program. Global variables remain in memory throughout the lifetime of the program. For example:
    
    ```go
    var x int // x has global scope
    
    func someFunction() {
        x = 10
    }
    
    func main() {
        fmt.Println(x) // prints 0
        someFunction()
        fmt.Println(x) // prints 10
    }
    ```
    


- When a variable with the same name is declared in a nested scope, it "shadows" the variable in the outer scope, meaning that the inner variable takes precedence over the outer one within that scope.

    > ⭐ It's important to note that if a variable is declared with the same name in both local and global scope, <span style="color:gold">the local variable takes precedence over the global variable</span> within the scope of the function or block where it is declared.

- Understanding variable scope is important for avoiding naming conflicts and for organizing code in a clear and structured way. It's generally recommended to keep the scope of variables as small as possible to avoid unintended side effects and make the code easier to understand and maintain.