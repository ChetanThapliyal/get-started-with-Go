package main

import "fmt"

func main() {

	//using var keyword
	// var variable_name data_type = value
	var age int = 26

	//shorthand method
	name := "John Doe"

	//var block
	var (
		company string = "Google"
		salary  int    = 26000
	)

	fmt.Println(company)
	fmt.Println(salary)
	fmt.Println(age)
	fmt.Println(name)
}
