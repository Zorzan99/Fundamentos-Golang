package formas

import (
	"math"
	"testing"
)

// Função de teste para verificar os cálculos de área
func TestArea(t *testing.T) {
	// Teste para o retângulo
	t.Run("Retângulo", func(t *testing.T) {
		// Cria um retângulo com largura 10 e altura 12
		ret := Retangulo{10, 12}
		// Área esperada do retângulo (10 * 12 = 120)
		areaEsperada := float64(120)
		// Calcula a área do retângulo
		areaRecebida := ret.Area()

		// Verifica se a área recebida é igual à área esperada
		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	// Teste para o círculo
	t.Run("Círculo", func(t *testing.T) {
		// Cria um círculo com raio 10
		circ := Circulo{10}
		// Área esperada do círculo (Pi * raio * raio)
		areaEsperada := float64(math.Pi * 100)
		// Calcula a área do círculo
		areaRecebida := circ.Area()

		// Verifica se a área recebida é igual à área esperada
		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
