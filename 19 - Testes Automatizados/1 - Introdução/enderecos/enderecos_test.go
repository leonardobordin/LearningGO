package enderecos

import (
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Avelinos"

	tipoDeEnderecoEsperado := "Rua"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
		t.Errorf("O tipo de endereço espera é diferente do recebido. Esperado: %s. Recebido: %s.",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido)
	}
}
