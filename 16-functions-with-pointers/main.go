package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1
}

func inverterSinalComPonteiro(num *int) {
	*num = *num * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)

	fmt.Println("numeroInvertido", numeroInvertido)
	fmt.Println("numero", numero)

	inverterSinalComPonteiro(&numero)
	fmt.Println("numero", numero)

	inverterSinalComPonteiro(&numero)
	fmt.Println("numero", numero)

	inverterSinalComPonteiro(&numero)
	fmt.Println("numero", numero)
}
