package main

import "fmt"

func main() {
	// bool
	var flag bool = true
	var zeroFlag bool
	fmt.Printf("bool: %v, zero value: %v\n", flag, zeroFlag)  //bool: true, zero value: false

	// numeric types
	var i8 int8 = -128
	var i16 int16 = -32768
	var i32 int32 = -2147483648
	var i64 int64 = -9223372036854775808
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i

	fmt.Printf("int8: %v\n", i8)  //int8: -128
	fmt.Printf("int16: %v\n", i16) //int16: -32768
	fmt.Printf("int32: %v\n", i32) //int32: -2147483648
	fmt.Printf("int64: %v\n", i64)  //int64: -9223372036854775808
	fmt.Printf("uint8: %v\n", u8)  //uint8: 255
	fmt.Printf("uint16: %v\n", u16)  //uint16: 65535
	fmt.Printf("uint32: %v\n", u32)  //uint32: 4294967295
	fmt.Printf("uint64: %v\n", u64)  //uint64: 18446744073709551615
	fmt.Printf("float32: %v\n", f32)  //float32: 3.14
	fmt.Printf("float64: %v\n", f64)  //float64: 3.141592653589793
	fmt.Printf("complex64: %v\n", c64)  //complex64: (1+2i)
	fmt.Printf("complex128: %v\n", c128)  //complex128: (3+4i)

	// string
	var str string = "Hello, World!"
	var emptyStr string
	fmt.Printf("string: %v, zero value: %v\n", str, emptyStr)  //string: Hello, World!, zero value:

	// arrays
	var arr [3]int = [3]int{1, 2, 3}
	var zeroArr [3]int
	fmt.Printf("array: %v, zero value: %v\n", arr, zeroArr)  //array: [1 2 3], zero value: [0 0 0]

	// slices
	var slice []int = []int{1, 2, 3}
	var emptySlice []int
	fmt.Printf("slice: %v, zero value: %v\n", slice, emptySlice)  //slice: [1 2 3], zero value: []

	// maps
	var m map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	var emptyMap map[string]int
	fmt.Printf("map: %v, zero value: %v\n", m, emptyMap)  //map: map[one:1 three:3 two:2], zero value: map[]


	// structs
	type person struct {
		name string
		age  int
	}
	var p person = person{"John", 30}
	var emptyPerson person
	fmt.Printf("struct: %v, zero value: %v\n", p, emptyPerson)  //struct: {John 30}, zero value: { 0}

}

