package main

import (
	"fmt"
	"math"
)

// interfaces só podem ter assinaturas de métodos
type Forma interface {
	area() float64
}

// function que recebe uma interface Forma
func EscreverArea(f Forma) {
	fmt.Printf("➡️  A área da forma é: %0.2f ⬅️\n", f.area())
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := Retangulo{altura: 10, largura: 15}
	EscreverArea(r)

	c := Circulo{raio: 10}
	EscreverArea(c)
}
