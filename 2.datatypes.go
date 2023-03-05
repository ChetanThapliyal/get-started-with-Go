package main
import "fmt"

func main() {
	//different types of variables

	//int
	var a int = 10
	var n int
	fmt.Println("Integers:")
	fmt.Println("int:", a)
	//zero value of datatype
	fmt.Println("Zero value of integer is:", n)
	
	//float
	var b float64 = 10.55
	var o float64
	fmt.Println("\n Float:")
	fmt.Println("float64:", b)
	//zero value of datatype
	fmt.Println("Zero value of float is:", o)

	//Boolean
	var c bool = true
	var p bool
	fmt.Println("\n Boolean:")
	fmt.Println("boolean:", c)
	//zero value of datatype
	fmt.Println("Zero value of boolean is:", p)

	//String
	var d string = "Oh My God"
	var q string
	fmt.Println("\n String:")
	fmt.Println("string:", d)
	//zero value of datatype
	fmt.Println("Zero value of string is:", q)

	//Array
	var e [3]int = [3]int {1,2,3}
	var r []int
	fmt.Println("\n Array:")
	fmt.Println("array:", e)
	//zero value of datatype
	fmt.Println("Zero value of array is:", r)

	//Slice
	var f []int = []int {1,2,3}
	var s []int
	fmt.Println("\n Slice:")
	fmt.Println("slice:", f)
	//zero value of datatype
	fmt.Println("Zero value of slice is:", s)

	//Map
	g := make(map[int]string)
	g[1] = "Ram"
	g[2] = "Laxman"
	g[3] = "Bharat"
	g[4] = "Shatrughan"
	var t map[int]string
	fmt.Println("\n Map:")
	fmt.Println("map:", g)
	//zero value of datatype
	fmt.Println("Zero value of map is:", t)


}
