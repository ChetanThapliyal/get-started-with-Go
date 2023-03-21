package main

import "fmt"

func main() {
    // create a slice of integers
    numbers := []int{10, 15, 20, 25, 30}

    // use a for loop to iterate over the slice
    for i := 0; i < len(numbers); i++ {
        // use an if-else statement to check if the current number is even or odd
        if numbers[i]%2 == 0 {
            fmt.Println(numbers[i], "is even")
        } else {
            fmt.Println(numbers[i], "is odd")
        }

        // use a switch statement to check the remainder of the current number divided by 5
        switch numbers[i] % 5 {
        case 0:
            fmt.Println(numbers[i], "is divisible by 5")
        case 1:
            fmt.Println(numbers[i], "has a remainder of 1 when divided by 5")
        case 2:
            fmt.Println(numbers[i], "has a remainder of 2 when divided by 5")
        case 3:
            fmt.Println(numbers[i], "has a remainder of 3 when divided by 5")
        default:
            fmt.Println(numbers[i], "has a remainder of 4 when divided by 5")
        }

        // use a while loop to repeatedly subtract 1 from the current number until it reaches 0
        for numbers[i] > 0 {
            // use a break statement to exit the loop if the current number is 10
            if numbers[i] == 10 {
                fmt.Println("Exiting the loop because", numbers[i], "was found")
                break
            }
            numbers[i]--
            fmt.Println("Current value:", numbers[i])
        }

        // use a continue statement to skip printing the message for the last number in the slice
        if i == len(numbers)-1 {
            continue
        }
        fmt.Println("Moving on to the next number...")
    }
}

