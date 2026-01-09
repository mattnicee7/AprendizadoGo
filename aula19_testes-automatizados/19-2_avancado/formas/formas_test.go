package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	// TDD - Test Driven Development
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f eh diferente da esperada %f", areaRecebida, areaEsperada)
			// Fatal = mesma coisa que o erro porém ele vai parar o teste aqui
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f eh diferente da esperada %f", areaRecebida, areaEsperada)
			// Fatal = mesma coisa que o erro porém ele vai parar o teste aqui
		}
	})
}
