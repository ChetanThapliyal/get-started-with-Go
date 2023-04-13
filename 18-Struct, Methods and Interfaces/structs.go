package main

import "fmt"

// Define a Person struct with two fields: Name and Age
type Person struct {
    Name string
    Age  int
}

// Define a function that takes a Person struct and prints out a greeting
func SayHello(p Person) {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}


// Define an UpdateAge function that takes a pointer to a Person struct and updates its Age field
func UpdateAge(p *Person, newAge int) {
    p.Age = newAge
}

func main() {
    // Initializing a struct using literal syntax
    person1 := Person{Name: "Alice", Age: 30}

    // Initializing a struct using anonymous syntax
    person2 := struct {
        Name string
        Age  int
    }{
        Name: "Bob",
        Age:  40,
    }

    // Initializing a struct using new function
    person3 := new(Person)
    person3.Name = "Charlie"
    person3.Age = 50

    // Initializing a struct with zero values
    var person4 Person

    // Accessing struct fields
    fmt.Println(person1.Name) // Output: Alice
    fmt.Println(person2.Age)  // Output: 40

    // Passing a struct to a function
    UpdateAge(&person1, 35)
    fmt.Println(person1.Age) // Output: 35

    // Comparing structs
    if person1 == person3 {
        fmt.Println("person1 and person3 are equal")
    } else {
        fmt.Println("person1 and person3 are not equal")
    }
}
