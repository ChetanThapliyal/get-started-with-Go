# Struct

A struct is a user-defined composite data type that groups together zero or more values of different types. Structs are used to represent data structures and can contain both named and anonymous fields.

Here's an example of a struct definition in Go:

```go
type Person struct {
    Name string
    Age  int
}
```

In this example, `Person` is a struct that contains two fields: `Name` of type `string` and `Age` of type `int`.

## Initializing Struct

There are several ways to initialize a struct.

### 1. Literal syntax:

The most common way to initialize a struct is using literal syntax, which involves specifying the values of the struct's fields in curly braces.

```go
type Person struct {
    Name string
    Age  int
}

person := Person{Name: "Alice", Age: 30}
```

In this example, a `Person` struct is created with `Name` set to "Alice" and `Age` set to 30.

### 2. Anonymous struct:

An anonymous struct is a struct that is defined without a name. It can be useful when you need to create a temporary struct with a specific set of values.

```go
person := struct {
    Name string
    Age  int
}{
    Name: "Bob",
    Age:  40,
}
```

In this example, an anonymous struct is created with `Name` set to "Bob" and `Age` set to 40.

### 3. Struct without Fields

You can create an instance of a struct without using the field names, as long as you define the fields in order:

```
s := Shape {
	"Oval",
	20,
}
```



### 4. New keyword:

The `new` keyword is used to allocate memory for all the fields of the new struct and returns a pointer to the newly created struct.

```go
type Person struct {
    Name string
    Age  int
}

personPtr := new(Person)
personPtr.Name = "Charlie"
personPtr.Age = 50
```

In this example, a new `Person` struct is created using `new`, and its fields are set individually using dot notation.

### 5. Zero value:

In Go, if you create a struct without initializing any of its fields, they will be set to their zero values. For strings, the zero value is `""`, and for integers, the zero value is `0`.

```go
type Person struct {
    Name string
    Age  int
}

func main() {
        var person Person	
        fmt.Printf("%+v", person) //%+v is used to print the contents of a struct's fields and its values, commonly used for debugging purposes.
}
```

In this example, a new `Person` struct is created with `Name` set to `""` and `Age` set to `0`.

### 6. **`&`** Operator:

You can create a new struct instance using the `&` operator, which returns a pointer to a newly created struct instance. Here's an example:

```go
type Person struct {
    Name string
    Age  int
}

// Using the & operator to initialize a struct
personPtr := &Person{Name: "Charlie", Age: 20}
```


> ⭐ Note that you can also use a combination of these methods to initialize a struct. For example, you could use literal syntax to initialize some fields, and then set others individually using dot notation.

## Accessing Struct Fields

You can access the fields of a struct using the dot notation (`.`). Here's an example:

```go
type Person struct {
    Name string
    Age  int
}

person1 := Person{Name: "Alice", Age: 30}

// Accessing struct fields using dot notation
fmt.Println(person1.Name) // Output: Alice
fmt.Println(person1.Age)  // Output: 30
```

In this example, we define a struct `Person` with two fields, `Name` and `Age`. We create a new `Person` instance called `person1` and initialize its fields.

To access the fields of the `person1` instance, we use the dot notation and print the values of `Name` and `Age` using the `fmt.Println` function.

You can also access the fields of a struct using a pointer to the struct. In this case, you use the dereference operator (`*`) to access the struct fields.

Here's an example:

```go
type Person struct {
    Name string
    Age  int
}

person1 := &Person{Name: "Alice", Age: 30}

// Accessing struct fields using a pointer
fmt.Println((*person1).Name) // Output: Alice
fmt.Println((*person1).Age)  // Output: 30
```

In this example, we create a pointer to a `Person` struct called `person1`. We use the dereference operator (`*`) to access the struct fields and print the values of `Name` and `Age` using the `fmt.Println` function.

## Passing Struct to Functions

In Go, you can pass structs to functions just like any other type.

Here's an example:

```go
type Person struct {
    Name string
    Age  int
}

func PrintPersonInfo(p Person) {
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

person1 := Person{Name: "Alice", Age: 30}

// Passing a struct to a function
PrintPersonInfo(person1)
```

In this example, we define a struct `Person` with two fields, `Name` and `Age`. We create a new `Person` instance called `person1` and initialize its fields.

We define a function `PrintPersonInfo` that takes a `Person` struct as a parameter and prints its fields.

We then call the `PrintPersonInfo` function and pass it the `person1` instance.

> When you pass a struct to a function, a copy of the struct is created and passed to the function. Any changes made to the struct within the function are local to that function and do not affect the original struct.

If you want to modify the original struct inside the function, you can pass a pointer to the struct instead.

Here's an example:

```go
type Person struct {
    Name string
    Age  int
}

func UpdatePersonAge(p *Person, newAge int) {
    p.Age = newAge
}

person1 := Person{Name: "Alice", Age: 30}

// Passing a pointer to a struct to a function
UpdatePersonAge(&person1, 35)

fmt.Printf("Name: %s, Age: %d\n", person1.Name, person1.Age) // Output: Name: Alice, Age: 35
```

In this example, we define a struct `Person` with two fields, `Name` and `Age`. We create a new `Person` instance called `person1` and initialize its fields.

We define a function `UpdatePersonAge` that takes a pointer to a `Person` struct and an integer `newAge` as parameters. We update the `Age` field of the struct using the pointer.

We then call the `UpdatePersonAge` function and pass it a pointer to the `person1` instance. After the function call, we print the `Name` and `Age` fields of the `person1` instance to verify that the `Age` field has been updated.

## Comparing Structs

In Go, you can compare two structs using the `==` operator. The `==` operator compares each field of the two structs for equality.

Here's an example:

```go
type Person struct {
    Name string
    Age  int
}

person1 := Person{Name: "Alice", Age: 30}
person2 := Person{Name: "Alice", Age: 30}
person3 := Person{Name: "Bob", Age: 25}

// Comparing two structs for equality
fmt.Println(person1 == person2) // Output: true
fmt.Println(person1 == person3) // Output: false
```

In this example, we define a struct `Person` with two fields, `Name` and `Age`. We create three new `Person` instances and initialize their fields.

We use the `==` operator to compare `person1` and `person2` for equality. Since the fields of the two structs have the same values, the comparison returns `true`.

We use the `==` operator again to compare `person1` and `person3` for equality. Since the `Name` and `Age` fields of the two structs have different values, the comparison returns `false`.


> ⭐ It's important to note that two structs with the same fields and values are not necessarily equal if they are of different types. In Go, each struct type is distinct and cannot be compared to another struct type, even if the two types have the same fields and values.

Here's an example:

```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Name string
    Age  int
}

person1 := Person{Name: "Alice", Age: 30}
employee1 := Employee{Name: "Alice", Age: 30}

// Comparing two structs of different types
fmt.Println(person1 == employee1) // Compile error: cannot compare Person and Employee
```

In this example, we define two struct types, `Person` and `Employee`, with the same fields and types. We create a `Person` instance called `person1` and an `Employee` instance called `employee1` with the same field values.

When we try to compare `person1` and `employee1` for equality using the `==` operator, we get a compile error because the two struct types are different and cannot be compared.

---

# Methods

A method is a function that is associated with a specific type of struct. It can be defined using the `func` keyword, the name of the method, the name of the struct type (with the `receiver` keyword), and any parameters and return types.

Whenever there's a strong relationship between a function and a struct, it makes sense to use a method.

Syntax :  **`func (<receiver>) <method_name>(<parameters>)<return_param> { //code}`**

Here's an example of a method definition in Go:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// A method associated with the Person struct
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I am %d years old\n", p.Name, p.Age)
}

func main() {
    // Creating a new Person instance
    person := Person{Name: "Alice", Age: 30}

    // Calling the SayHello method on the Person instance
    person.SayHello() //Output: Hello, my name is Alice and I am 30 years old
}
```

In this example, we define a `Person` struct with two fields, `Name` and `Age`. We also define a method `SayHello` associated with the `Person` struct.

The `SayHello` method has a `receiver` of type `Person`, which means that it is associated with the `Person` struct. The `receiver` is passed as the first argument to the method and is accessed using the `p` variable.

Inside the `SayHello` method, we use the `fmt.Printf` function to print a message that includes the `Name` and `Age` fields of the `Person` instance.

In the `main` function, we create a new `Person` instance called `person` and initialize its fields. We then call the `SayHello` method on the `person` instance using the `.` operator.

This example demonstrates how to define a method associated with a struct type, and how to call the method on an instance of the struct using the `.` operator.

It's important to note that methods can also be defined with pointers to the struct as the receiver. This allows the method to modify the struct's fields. Here's an example:

```go
// A method associated with the Person struct using a pointer receiver
func (p *Person) SetAge(age int) {
    p.Age = age
}
```

This method takes a pointer to the `Person` struct as its receiver (`*Person`), which allows it to modify the `Age` field of the struct.

## Method Sets

A method set is the set of methods that can be called on a specific type of struct.

There are two types of method sets in Go:

1. Value receiver method set
2. Pointer receiver method set

The value receiver method set contains all of the methods that are associated with a struct using a value receiver. These methods can be called on an instance of the struct or a pointer to the struct.

The pointer receiver method set contains all of the methods that are associated with a struct using a pointer receiver. These methods can only be called on a pointer to the struct.

Here's an example that demonstrates how method sets work:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// A value receiver method associated with the Person struct
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I am %d years old\n", p.Name, p.Age)
}

// A pointer receiver method associated with the Person struct
func (p *Person) SetAge(age int) {
    p.Age = age
}

func main() {
    // Creating a new Person instance
    person := Person{Name: "Alice", Age: 30}

    // Calling the SayHello method on the Person instance
    person.SayHello()

    // Calling the SetAge method on the Person instance using a pointer
    (&person).SetAge(35)

    // Calling the SayHello method on the Person instance again
    person.SayHello()

    // Creating a pointer to the Person instance
    pointer := &person

    // Calling the SetAge method on the pointer to the Person instance
    pointer.SetAge(40)

    // Calling the SayHello method on the Person instance again
    person.SayHello()
}
```

In this example, we define a `Person` struct with two fields, `Name` and `Age`. We also define two methods, `SayHello` and `SetAge`, associated with the `Person` struct.

The `SayHello` method is associated with the `Person` struct using a value receiver. This means that it can be called on an instance of the `Person` struct or a pointer to the struct.

The `SetAge` method is associated with the `Person` struct using a pointer receiver. This means that it can only be called on a pointer to the `Person` struct.

In the `main` function, we create a new `Person` instance called `person` and initialize its fields. We then call the `SayHello` method on the `person` instance.

We then call the `SetAge` method on the `person` instance using a pointer. This works because we are calling a method from the value receiver method set on a pointer to the `Person` struct.

We then create a pointer to the `person` instance called `pointer`. We call the `SetAge` method on the `pointer` to the `Person` instance. This works because we are calling a method from the pointer receiver method set on a pointer to the `Person` struct.

Finally, we call the `SayHello` method on the `person` instance again to see the updated value of the `Age` field.

Output:

```go
Hello, my name is Alice and I am 30 years old
Hello, my name is Alice and I am 35 years old
Hello, my name is Alice and I am 40 years old
```
---

# Interfaces

Interfaces are a set of method signatures that a type must implement in order to satisfy the interface. Interfaces provide a way to specify behavior without specifying the underlying implementation. This makes interfaces a powerful tool for decoupling code and promoting code reuse. 

A type implements an interface by implementing its methods. The go language interfaces are implemented implicitly. And it does not have any specific keyword to implement an interface.

Here's an example that demonstrates how to define and use interfaces in Go:

```go
package main

import "fmt"

// Define the Shape interface
type Shape interface {
    Area() float64
}

// Define the Circle struct
type Circle struct {
    x, y, r float64
}

// Define the Area method on the Circle struct
func (c Circle) Area() float64 {
    return 3.14 * c.r * c.r
}

// Define the Rectangle struct
type Rectangle struct {
    width, height float64
}

// Define the Area method on the Rectangle struct
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
    // Create a Circle instance
    circle := Circle{x: 0, y: 0, r: 5}

    // Create a Rectangle instance
    rectangle := Rectangle{width: 10, height: 5}

    // Define a slice of Shape interfaces and add Circle and Rectangle instances to it
    shapes := []Shape{circle, rectangle}

    // Loop through the shapes slice and call the Area method on each instance
    for _, shape := range shapes {
        fmt.Printf("Area of shape: %.2f\n", shape.Area())
    }
}
```

In this example, we define an interface called `Shape` with a single method signature called `Area`, which returns a float64. We then define two structs, `Circle` and `Rectangle`, and implement the `Area` method on each struct.

In the `main` function, we create a `Circle` instance and a `Rectangle` instance. We then define a slice of `Shape` interfaces and add the `Circle` and `Rectangle` instances to it. This works because both `Circle` and `Rectangle` implement the `Area` method defined by the `Shape` interface.

We then loop through the `shapes` slice and call the `Area` method on each instance. Because both `Circle` and `Rectangle` implement the `Area` method, the correct implementation is called for each instance.

The output of the program will be:

```
Area of shape: 78.50
Area of shape: 50.00
```

This example demonstrates how to define and use interfaces in Go. By defining an interface with a set of method signatures, we can create types that implement the interface and use them interchangeably.