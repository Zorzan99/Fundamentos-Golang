package main

import "fmt"

func main() {
	// Declaração e inicialização de variáveis
	var variavel1 string = "Variavel"
	variavel2 := "Variavel 2"

	// Impressão de variáveis
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declaração e inicialização de múltiplas variáveis usando o bloco var
	var (
		variavel3 string = "Teste3"
		variavel4 string = "Teste4"
	)
	fmt.Println(variavel3, variavel4)

	// Declaração e inicialização de múltiplas variáveis de forma concisa
	variavel5, variavel6 := "Variavel 5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	// Declaração de constante
	const constante1 string = "Constante1"
	fmt.Println(constante1)

	// Troca de valores entre variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}

// Explicação:

// 1 package main: Define que este é um programa executável em Go e pertence ao pacote principal.

// 2 import "fmt": Importa o pacote fmt, que fornece funcionalidades de formatação de texto para impressão na saída padrão.

// 3  var variavel1 string = "Variavel": Declaração e inicialização de uma variável chamada variavel1 do tipo string.

// 4 variavel2 := "Variavel 2": Declaração e inicialização de uma variável variavel2 de forma concisa. O tipo é inferido automaticamente.

// 5 fmt.Println(variavel1): Imprime o valor da variavel1 usando a função Println do pacote fmt.

// 6 var (...): Bloco de declaração de várias variáveis.

// 7 variavel5, variavel6 := "Variavel 5", "Variavel6": Declaração e inicialização de duas variáveis (variavel5 e variavel6) de forma concisa.

// 8 const constante1 string = "Constante1": Declaração de uma constante chamada constante1 do tipo string.

// 9 variavel5, variavel6 = variavel6, variavel5: Troca os valores das variáveis variavel5 e variavel6.

// 10 O programa imprime os valores das variáveis e da constante usando fmt.Println.
