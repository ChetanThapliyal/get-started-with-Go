# GO Fundamentals

Go programs are organized into packages. Every Go file starts with a package clause.

To run the program:

```bash
#To run the program, put the code in hello-world.go and use go run.
go run hello-world.go
hello world

##Sometimes we’ll want to build our programs into binaries. 
##We can do this using go build.
go build hello-world.go

ls
hello-world    hello-world.go	

## We can then execute the built binary directly.
./hello-world
hello world
```

## Typical GO Layout

In almost every Go file you work with:

1. The package clause
2. Any import statements
3. The actual code

## Package
A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

- The package clause gives the name of the package that this file’s code will become a part of.
- In “Hello World” example, we use the special package **main**, which is required if this code is going to be run directly (usually from the terminal).

## Module 
A repository contains one or more modules. A module is a collection of related Go packages that are released together.

- A Go repository typically contains only one module, located at the root of the repository. 
- A file named `go.mod` there declares the module path: the import path prefix for all packages within the module. 
- The module contains the packages in the directory containing its `go.mod` file as well as subdirectories of that directory, up to the next subdirectory containing another `go.mod` file (if any).

> 📌 **Note** that you don't need to publish your code to a remote repository before you can build it. A module can be defined locally without belonging to a repository.

- Each module's path not only serves as an import path prefix for its packages, but also indicates where the `go` command should look to download it.

## Import
Go files almost always have one or more import statements. An import path is a string used to import a package. 

- Each file needs to import other packages before its code can use the code those other packages contain. Loading all the Go code on your computer at once would result in a big, slow program, so instead you specify only the packages you need by importing them.
- A package's import path is its module path joined with its subdirectory within the module. For example, the module [go-cmp](https://github.com/google/go-cmp) contains a package in the directory `cmp/`. That package's import path is <https://github.com/google/go-cmp/cmp>. Packages in the standard library do not have a module path prefix.

**Function**: The last part of every Go file is the actual code, which is often split up into one or more functions. A function is a group of one or more lines of code that you can call (run) from other places in your program. 

- When a Go program is executed, it looks for a function named main and runs that first, which is why we named this function main. (acts as an entry point for the compiler.)