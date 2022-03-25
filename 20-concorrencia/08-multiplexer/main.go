package main

import (
	"fmt"
	"math/rand"
	"time"
)

func escrever(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplex(canal1 <-chan string, canal2 <-chan string) <-chan string {
	result := make(chan string)

	go func() {
		for {
			select {
			case msg := <-canal1:
				result <- msg
			case msg := <-canal2:
				result <- msg
			}
		}
	}()

	return result
}

func main() {
	canal := multiplex(escrever("hello world"), escrever("ola mundo"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
