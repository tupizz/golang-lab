package main

import "fmt"

func recoverFromPanic() {
	fmt.Println("trying to recover execution")

	if r := recover(); r != nil {
		fmt.Println("function recovered successfully")
	}
}

func isStudentApproved(scoreSlice ...int) bool {
	defer recoverFromPanic()

	scores := 0
	for _, score := range scoreSlice {
		scores += score
	}

	mean := scores / len(scoreSlice)

	if mean == 6 {
		panic("A MEDIA NÃO PODE SER 6")
	}

	return mean > 6
}

func main() {
	// we treat panic like this
	isApproved := isStudentApproved(6, 6, 6, 6, 6)
	fmt.Println(isApproved)
}
