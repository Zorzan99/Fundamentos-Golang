package enderecos

// Importação do pacote "strings" para manipulação de strings
import "strings"

// Função TipoDeEndereco recebe um endereço como parâmetro e retorna o tipo de endereço
func TipoDeEndereco(endereco string) string {
	// Lista de tipos de endereço válidos
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	// Converte o endereço para letras minúsculas para tornar a comparação case-insensitive
	enderecoComLetraMinuscula := strings.ToLower(endereco)

	// Obtém a primeira palavra do endereço usando espaço como delimitador
	primeiraPalavraDoEndereco := strings.Split(enderecoComLetraMinuscula, " ")[0]

	// Verifica se a primeira palavra do endereço está na lista de tipos válidos
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	// Se o endereço tiver um tipo válido, retorna o tipo com a primeira letra maiúscula; caso contrário, retorna "Tipo inválido"
	if enderecoTemUmTipoValido {
		return strings.ToUpper(primeiraPalavraDoEndereco[:1]) + primeiraPalavraDoEndereco[1:]
	}

	return "Tipo inválido"
}
