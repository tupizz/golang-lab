package main

import "fmt"

func main() {
	// int8, int16, int32, int64
	// int8 = 2^8 = 256 = -128 to +(128 - 1)
	// int16 = 2^16 = 65536 = -32768 to +(32768 - 1)
	// (...)
	// int -> get 32 or 64 based on the architecture of your processor
	// int32 = rune
	// int8 = byte

	var num int8 = 127
	fmt.Println(num)
	fmt.Println(num + 1)

}
