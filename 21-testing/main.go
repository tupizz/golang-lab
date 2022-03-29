package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	addressType := enderecos.GetAddressType("Rua Maria Benedita de Oliveira")
	fmt.Println(addressType)
}
