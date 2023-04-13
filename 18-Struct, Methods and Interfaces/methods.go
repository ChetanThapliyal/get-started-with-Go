package main

import "fmt"

type Rectangle struct {
    width, height int
}

// A method associated with the Rectangle struct using a value receiver
func (r Rectangle) Area() int {
    return r.width * r.height
}

// A method associated with the Rectangle struct using a pointer receiver
func (r *Rectangle) Scale(factor int) {
    r.width *= factor
    r.height *= factor
}

func main() {
    // Initializing a Rectangle struct using a struct literal
    rectangle := Rectangle{width: 10, height: 5}

    // Calling the Area method on the Rectangle instance
    area := rectangle.Area()
    fmt.Printf("Area of rectangle: %d\n", area)

    // Calling the Scale method on the Rectangle instance using a pointer
    rectangle.Scale(2)
    fmt.Printf("Scaled rectangle: %+v\n", rectangle)

    // Creating a pointer to the Rectangle instance
    pointer := &rectangle

    // Calling the Scale method on the pointer to the Rectangle instance
    pointer.Scale(3)
    fmt.Printf("Scaled rectangle: %+v\n", rectangle)
}
