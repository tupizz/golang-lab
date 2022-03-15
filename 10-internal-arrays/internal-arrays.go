package main

import "fmt"

func main() {
	slice := make([]float32, 10, 12) // por trás dos panos criamos um array de 12 posições, mas retornamos um slice de length igual a 10 para manipular
	fmt.Println(slice)
	fmt.Println("tamanho:", len(slice), " capacidade: ", cap(slice))

	// aqui estamos adicionando mais dois items no slice de 10 sobre o array interno de 12
	slice = append(slice, 10.3)
	slice = append(slice, 10.5)
	fmt.Println(slice)
	fmt.Println("tamanho:", len(slice), " capacidade: ", cap(slice))

	// Resultado:
	// [0 0 0 0 0 0 0 0 0 0 10.3 10.5]
	// tamanho: 12  capacidade:  12

	// Aqui estamos adicionando mais um float value no slice de capacidade 12, estourando a capacidade do slice
	slice = append(slice, 13.3)
	fmt.Println(slice)
	fmt.Println("tamanho:", len(slice), " capacidade: ", cap(slice))

	// Resultado:
	// [0 0 0 0 0 0 0 0 0 0 10.3 10.5 13.3]
	// tamanho: 13  capacidade:  24
}
