package main

import (
	"log"
	"reflect"
	"runtime"
	"time"
)

func calculos(n1, n2 int) (soma int, multiplicacao int) {
	soma = n1 + n2
	multiplicacao = n1 * n2
	return
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func createMemoizeFib() func(uint) uint {
	cache := map[uint]uint{}

	var fibFn func(uint) uint
	fibFn = func(posicao uint) uint {
		if posicao <= 1 {
			return posicao
		}

		if _, ok := cache[posicao]; !ok {
			cache[posicao] = fibFn(posicao-1) + fibFn(posicao-2)
		}

		return cache[posicao]
	}

	return fibFn
}

func timeWrapper(function func(uint) uint, param uint) {
	start := time.Now()
	function(param)
	elapsed := time.Since(start)
	log.Printf("Function %s took %s", runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name(), elapsed)
}

func main() {
	func() {
		fibFunction := createMemoizeFib()
		timeWrapper(fibFunction, 40)
		timeWrapper(fibonacci, 40)
	}()

}
