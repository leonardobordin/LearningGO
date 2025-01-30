package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposDeEndereco := []string{"rua", "avenida", "rodavia", "bairro"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemTipoValido := false
	for _, tipo := range tiposDeEndereco {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
