package formas_test

import (
	. "advanced-tests/formas"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("a área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		c := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := c.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("a área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
