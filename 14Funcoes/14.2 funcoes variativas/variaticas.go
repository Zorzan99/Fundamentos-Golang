package main

import "fmt"

// Função que realiza a soma de uma quantidade variável de números inteiros
func soma(numeros ...int) int {
	total := 0

	// Loop usando range para iterar sobre os números passados como argumentos
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Função que imprime um texto seguido de uma lista de números inteiros
func escrever(texto string, numeros ...int) {
	// Loop usando range para iterar sobre os números passados como argumentos
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// Chamada da função soma com uma lista variável de argumentos
	totalDasoma := soma(1, 2, 3, 4, 5, 6)
	fmt.Println(totalDasoma)

	// Chamada da função escrever com um texto e uma lista variável de números
	escrever("Olá mundo", 1, 2)
}

//Neste código, são utilizadas funções com parâmetros variáveis (...) para realizar a soma de uma quantidade variável de números inteiros e imprimir um texto seguido de uma lista de números.
// Os comentários explicam cada parte do código
