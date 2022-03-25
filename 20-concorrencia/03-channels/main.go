package main

import (
	"fmt"
	"time"
)

func escrever(text string, canal chan string) {
	times := 10
	for i := 0; i < times; i++ {
		time.Sleep(time.Second / 2)
		canal <- text
	}

	close(canal)
}

// syncing with channel
func main() {
	// criando um canal que traféga strings
	canal := make(chan string)

	go escrever("olá mundo", canal)

	// for {
	// 	mensagem, aberto := <-canal

	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("fim do programa")
}
