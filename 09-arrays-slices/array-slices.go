package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("ARRAY AND SLICES")

	/**
	 * Arrays
	 */
	var array1 [5]int
	array1[0] = 10
	array1[1] = 12
	fmt.Println(array1)

	array2 := [6]string{"tadeu", "humberto", "dos", "reis", "tupinamba", "junior"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // use ... to determine the length of array given their init values
	fmt.Println(array3)

	/**
	 * Slices - is the most used across golang
	 */
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)
	fmt.Println("typeof slice", reflect.TypeOf(slice))  // []int
	fmt.Println("typeof array", reflect.TypeOf(array1)) // [5]int

	fmt.Println("adding items into slice")
	slice = append(slice, 16)
	fmt.Println(slice)
	slice = append(slice, 17)
	fmt.Println(slice)
	slice = append(slice, 18)
	fmt.Println(slice)

	sliceOfArray := array3[2:4] // pointer to sub array of items
	fmt.Println(sliceOfArray)
}
