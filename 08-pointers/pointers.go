package main

import "fmt"

func main() {
	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2)

	variable1++

	fmt.Println(variable1, variable2) // when we change the value from variable 1 it does not update variable 2

	// Pointer is a memory's reference
	var variable3 int
	var pointer *int

	variable3 = 10
	pointer = &variable3 // we are referencing the address memory of variable3

	fmt.Println(pointer, *pointer) // address memory and value --> 0x1400001c0a0 10
	variable3 = 150
	fmt.Println(pointer, *pointer) // address memory and value --> 0x1400001c0a0 10
}
