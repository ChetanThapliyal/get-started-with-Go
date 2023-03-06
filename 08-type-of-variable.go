package main
import (
	"fmt"
	"reflect"
)

func main () {
	var grades int = 32
	var cgpa float32 = 9.2
	var istrue bool = true
	var work string = "Good Work"


	// Using Prinf with %T
	fmt.Printf("variable grades = %v is of type %T \n", grades,grades)
	fmt.Printf("variable cgpa = %v is of type %T \n", cgpa, cgpa)
	fmt.Printf("variable istrue = %v is of type %T \n", istrue, istrue)
	fmt.Printf("variable work = %v is of type %T \n", work, work)

	// Using reflect package
	fmt.Printf("Type of variable grades is: %v \n", reflect.TypeOf(grades))
	fmt.Printf("Type  of variable istrue is: %v \n", reflect.TypeOf(istrue))
	fmt.Printf("Type  of variable cgpa is: %v \n", reflect.TypeOf(cgpa))
	fmt.Printf("Type  of variable work is: %v \n", reflect.TypeOf(work))

}
