package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	multiplicacao := n1 * n2
	return soma, multiplicacao
}

func main() {
	soma := somar(10, 2)

	fn := func(txt string) string {
		fmt.Println(txt + "oi")
		return txt + "oi"
	}

	fn("tadeu")
	fmt.Println(calculos(4, 5))
	fmt.Println(soma)
}
