package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i value", i)
		time.Sleep(time.Second)
	}

	nomes := []string{"tadeu", "humberto", "tupinamba"}
	for index, value := range nomes {
		fmt.Println(index, " -> ", value)
		time.Sleep(time.Second)
	}

	for index, charCode := range "TADEU" {
		fmt.Println(index, " -> ", string(charCode))
		time.Sleep(time.Second)
	}

}
