package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestGetAddressType(t *testing.T) {
	t.Parallel() // roda em paralelo com outros testes

	cenarioDeTeste := []cenarioDeTeste{
		{
			enderecoInserido: "Avenida Paulista",
			enderecoEsperado: "Avenida",
		},
		{
			enderecoInserido: "Avenida Pernambuco, 2746",
			enderecoEsperado: "Avenida",
		},
		{
			enderecoInserido: "Rua Maria Benedita de Oliveira",
			enderecoEsperado: "Rua",
		},
		{
			enderecoInserido: "Praça governador Valadares",
			enderecoEsperado: "Invalid type",
		},
	}

	for _, cenario := range cenarioDeTeste {
		result := GetAddressType(cenario.enderecoInserido)
		if result != cenario.enderecoEsperado {
			t.Errorf("Endereço recebido '%s' foi diferente to endereço inserido '%s'", result, cenario.enderecoInserido)
		}
	}
}
