package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		retanc := Retangulo{60, 10}

		areaEsperada := float64(600)
		areaRecebida := retanc.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Teste falhou. Area Esperada: %f. Area Recebida: %f", areaEsperada, areaRecebida)
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Teste falhou. Area Esperada: %f. Area Recebida: %f", areaEsperada, areaRecebida)
		}
	})
}
