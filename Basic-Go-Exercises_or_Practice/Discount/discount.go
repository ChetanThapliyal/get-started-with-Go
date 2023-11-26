package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	var discount float64

	switch product {
	case "apples":
		discount = 0.1
	case "bananas":
		discount = 0.2
	default:
		discount = 0
	}

}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
