package main

import "fmt"

type Person struct {
	name     string
	lastName string
	age      int8
}

type Student struct {
	Person // similar to extends
	course string
}

func main() {
	p1 := Person{name: "Joao", lastName: "alves", age: 27}
	fmt.Println(p1)

	s1 := Student{Person: p1, course: "engenharia"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.lastName)
}
