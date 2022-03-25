package main

import (
	"fmt"
	"sync"
	"time"
)

// wait group = sincronizar go routines

func escrever(text string) {
	times := 10
	for i := 0; i < times; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println(text)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	// duas go routines para esperar terminar
	waitGroup.Add(4)

	// utilizamos uma anonymous function
	go func() {
		escrever("olá mundo")
		// tira do contador
		waitGroup.Done()
	}()

	// utilizamos uma anonymous function
	go func() {
		escrever("Salut Lume")
		// tira do contador
		waitGroup.Done()
	}()

	// utilizamos uma anonymous function
	go func() {
		escrever("hello world")
		// tira do contador
		waitGroup.Done()
	}()

	// utilizamos uma anonymous function
	go func() {
		escrever("Hola mundo")
		// tira do contador
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
