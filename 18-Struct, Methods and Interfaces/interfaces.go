package main

import "fmt"

// Define the User interface
type User interface {
    GetName() string
}

// Define the User struct
type Student struct {
    name string
}

// Define the GetName method on the Student struct
func (s Student) GetName() string {
    return s.name
}

// Define the UserPrinter interface
type UserPrinter interface {
    PrintUser(User)
}

// Define the ConsoleUserPrinter struct
type ConsoleUserPrinter struct{}

// Define the PrintUser method on the ConsoleUserPrinter struct
func (p ConsoleUserPrinter) PrintUser(user User) {
    fmt.Printf("User: %s\n", user.GetName())
}

func main() {
    // Create a Student instance
    student := Student{name: "John Doe"}

    // Create a ConsoleUserPrinter instance
    printer := ConsoleUserPrinter{}

    // Call the PrintUser method on the printer instance, passing in the student instance
    printer.PrintUser(student)
}
