package enderecos

import (
	"strings"
)

func GetAddressType(endereco string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}
	firstWord := strings.Split(strings.ToLower(endereco), " ")[0]

	addressHasValidType := false

	for _, tipo := range validTypes {
		if tipo == firstWord {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return strings.Title(firstWord)
	}

	return "Invalid type"
}
