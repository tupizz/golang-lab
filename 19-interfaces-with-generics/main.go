package main

import "fmt"

func funcaoGenerica(interf interface{}) {
	fmt.Println(interf)
}

type Pessoa struct {
	nome  string
	idade int8
}

func main() {
	funcaoGenerica("tadeu")
	funcaoGenerica(10)
	funcaoGenerica(true)
	funcaoGenerica(Pessoa{nome: "tadeu", idade: 27})
}
