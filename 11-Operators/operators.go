package main

import "fmt"

func main() {
	// arithmetic operators
	a := 10
	b := 3
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 3
	fmt.Println(a % b) // 1
	a++
	fmt.Println(a) // 11
	b--
	fmt.Println(b) // 2

	// comparison operators
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false

	// logical operators
	x := true
	y := false
	fmt.Println(x && y) // false  (AND)
	fmt.Println(x || y) // true   (OR)
	fmt.Println(!x)     // false  (NOT)

	// assignment operators
	c := 5
	c += 2         // equivalent to c = c + 2
	fmt.Println(c) // 7
	c *= 3         // equivalent to c = c * 3
	fmt.Println(c) // 21
	c %= 4         // equivalent to c = c % 4
	fmt.Println(c) // 1

	// other operators
	p := &c              // p is a pointer to c
	fmt.Println(p)       // 0xc0000180a0 (memory address of c)
	fmt.Println(*p)      // 1 (value of c)
	ch := make(chan int) // create a channel
	go func() {
		ch <- 42 // send 42 on the channel
	}()
	fmt.Println(<-ch) // receive 42 from the channel
	res := ""
	if a > b {
		res = "a > b"
	} else {
		res = "a <= b"
	}
	fmt.Println(res) // a > b
}

