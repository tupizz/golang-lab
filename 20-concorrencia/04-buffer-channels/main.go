package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "ola mundo!"
	canal <- "hello world!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem, mensagem2)
}
