package main

import "fmt"

func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos-2) + fibonacci(pos-1)
}

// tarefas - só envia dados <-chan
// resultados só recebe dados chan<-
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

// crio um canal de entrada
// crio um canal de saida
// defino/replico a chamada do worker que vai criar mais um worker para ler dos input,
// processar e jogar no canal de resultados
func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
