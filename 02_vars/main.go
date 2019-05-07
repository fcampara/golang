package main

import "fmt"

var name = "Felipe"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// bye - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name string = "Felipe"

	// Shorthand only can in inside function
	// name := "Felipe"
	size := 1.3
	// email := "fcamparasilva@gmail.com"

	name, email := "Felipe Campara", "felipe_novo2@hotmail.com"

	var age int = 26
	const isCool bool = true

	fmt.Println(name, email, age, isCool, size)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
