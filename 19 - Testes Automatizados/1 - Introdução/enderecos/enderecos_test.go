package enderecos_test

import (
	. "introducao-testes/enderecos" // Passa . como alias para ele entender que o Enderecos é o padrão então não preciso fazer enderecos.TipoDeEndereco
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	tipoEsperado     string
}

func TestTipoDeEndereco(t *testing.T) {
	// t -> t.Parallel -> Roda os testes em paralelo
	// go test -v  -> Modo verboso será mais descritivo
	// go test --cover -> Este comando dirá qual a cobertura do teste no caso quantos cenarios ele está cobrindo
	// go test --coverprofile cobertura.txt -> salva em um arquivo o quais cenarios não estão sendo testados
	// go tool cover --func=cobertura.txt -> De acordo com o arquivo salvo mostra a função que não está sendo testada
	// go tool cover --html=cobertura.txt -> De acordo com o arquivo salvo mostra um pagina HMLT exatamente o passo que não foi testado
	t.Parallel()
	cenarios := []cenarioDeTeste{
		{"Rua Manoel Gomes", "Rua"},
		{"Avenida Antonio Alves", "Avenida"},
		{"Rodovia Luis Carlos", "Rodovia"},
		{"Bairro Jardim Bela Vista", "Bairro"},
		{"José Antonio", "Tipo Inválido"},
	}

	for _, cenario := range cenarios {
		tipoDeEndereco := TipoDeEndereco(cenario.enderecoInserido)
		if cenario.tipoEsperado != tipoDeEndereco {
			t.Errorf("O tipo de endereço difere do esperado. Esperado: %s. Recebido: %s.", cenario.tipoEsperado, tipoDeEndereco)
		}

	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste Quebrou!")
	}
}
