package main

import (
	"fmt"
	"time"
)

// timeTrack is a helper function that takes a start time and a name as arguments,
// and prints a message indicating how long it took to execute the function.
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

// expensiveFunction is a function that takes a long time to execute.
// We're using a defer statement to time how long it takes to execute.
func expensiveFunction() {
	defer timeTrack(time.Now(), "expensiveFunction")

	// Do something that takes a long time...
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}

func main() {
	fmt.Println("Starting...")
	expensiveFunction()
	fmt.Println("Finished!")
}