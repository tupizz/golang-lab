package main

import "fmt"

type Usuario struct {
	name  string
	age   int8
	email string
}

func (u Usuario) save() {
	fmt.Println("salvando user", u)
}

func (u *Usuario) makeBirthday() {
	u.age++
}

func createUser() Usuario {
	return Usuario{
		name:  "usuario",
		age:   34,
		email: "tadeu.tupiz@gmail.com",
	}
}

func main() {
	user := createUser()
	user.name = "new name"
	user.save()
	user.makeBirthday()
	user.save()
	user.makeBirthday()
	user.save()
}
