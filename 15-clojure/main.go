package main

import "fmt"

func clojure() func() {
	text := "dentro da funcao clojure"

	// function will remember where it was defined
	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "dentro da função main"
	fmt.Println(text)

	newfn := clojure()
	newfn() // dentro da funcao clojure
}
