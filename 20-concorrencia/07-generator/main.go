package main

import (
	"fmt"
	"time"
)

func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}

func main() {
	canal := escrever("ola mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
