package main

import "fmt"

var globalVar string = "I am a global variable"

func main() {
    localVar := "I am a local variable"
    fmt.Println(localVar)
    fmt.Println(globalVar)
    outerFunction()
}

func outerFunction() {
    localVar := "I am a local variable in outerFunction"
    fmt.Println(localVar)
    fmt.Println(globalVar)
    innerFunction()
}

func innerFunction() {
    localVar := "I am a local variable in innerFunction"
    fmt.Println(localVar)
    fmt.Println(globalVar)
}

