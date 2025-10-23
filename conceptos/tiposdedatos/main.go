package main

import "fmt"

func main() {
	//var a bool = true
	var a uint8 = 12
	var b uint16 = 3500

	c:= uint16(a) + b

	//fmt.Println(c)
	//fmt.Println(a)
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("%T, %v\n", a, a)	
}

// tipos de datos
// bool TRUE/FALSE
// string "Hola"
// int 123
// int8 -128 to 127
// int16 -32768 to 32767
// int32 -2147483648 to 2147483647
// int64 -9223372036854775808 to 9223372036854775807
// uint 0 to 255
// uint8 0 to 255
// uint16 0 to 65535
// uint32 0 to 4294967295
// uint64 0 to 18446744073709551615
// float32 123.12
// float64 123.12
// complex64 123.12
// complex128 123.12	