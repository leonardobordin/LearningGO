package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
)

func TipoDeEndereco(endereco string) string {
	tiposDeEndereco := []string{"rua", "avenida", "rodovia", "bairro"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemTipoValido := false
	for _, tipo := range tiposDeEndereco {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return cases.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
