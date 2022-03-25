package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	// espera meio segundo
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal 1"
		}
	}()

	// espera dois segundos
	go func() {
		for {
			time.Sleep(time.Second * 5)
			canal2 <- "canal 2"
		}
	}()

	for {
		// dessa forma abaixo o canal que demora mais

		select {
		case msgCanal1 := <-canal1:
			fmt.Println(msgCanal1)
		case msgCanal2 := <-canal2:
			fmt.Println(msgCanal2)
		}

		// dessa forma abaixo o canal que demora mais bloqueia o canal que é mais rapido

		// msgCanal1 := <-canal1
		// fmt.Println(msgCanal1)
		// msgCanal2 := <-canal2
		// fmt.Println(msgCanal2)
	}

}
