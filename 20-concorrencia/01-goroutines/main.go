package main

import (
	"fmt"
	"time"
)

// Concorrencia != paralelismo

func escrever(text string) {
	for {
		time.Sleep(time.Second)
		fmt.Println(text)
	}
}

func main() {
	go escrever("olá mundo")
	escrever("adeus mundo")
}
