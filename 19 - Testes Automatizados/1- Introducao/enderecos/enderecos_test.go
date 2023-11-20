// TESTE DE UNIDADE
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

// Definição de uma estrutura (struct) para representar um cenário de teste
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Função de teste para TipoDeEndereco
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	// Definição de uma lista de cenários de teste
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos imigrantes", "Rodovia"},
		{"Praça das rosas", "Tipo inválido"},
		{"Estrada qualquer", "Estrada"},
		{"Avenida Rebolças", "Avenida"},
		{"", "Tipo inválido"},
	}

	// Iteração sobre cada cenário de teste
	for _, cenario := range cenariosDeTeste {
		// Chama a função TipoDeEndereco com o endereço do cenário
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		// Compara o resultado recebido com o resultado esperado
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			// Emite um erro se os resultados não coincidirem
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}

// //A estrutura cenarioDeTeste é usada para representar cada cenário de teste, consistindo em um endereço inserido e o resultado esperado da função TipoDeEndereco.

// O loop for itera sobre cada cenário de teste na lista cenariosDeTeste.

// Para cada cenário de teste, a função TipoDeEndereco é chamada com o endereço fornecido, e o resultado é comparado com o resultado esperado.

// Se os resultados não coincidirem, um erro é emitido usando t.Errorf, indicando os valores recebidos e esperados.
