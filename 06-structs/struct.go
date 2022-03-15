package main

import "fmt"

type User struct {
	name  string
	idade int8
}

func main() {
	var usuario User
	usuario.idade = 19
	usuario.name = "joaozinho"
	fmt.Println(usuario)

	newUser := User{name: "tadeu", idade: 27}
	fmt.Println(newUser)

	users := []User{
		{name: "tadeu", idade: 27},
		{name: "henrique", idade: 27},
		{name: "gabriel", idade: 27},
	}

	fmt.Println(users)
}
