package main

import "fmt"

// Função recursiva para calcular o número Fibonacci na posição dada
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	// Chama recursivamente a função para as posições anteriores e soma os resultados
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// Posição desejada na sequência de Fibonacci
	posicao := uint(30)

	// Loop para imprimir os números Fibonacci até a posição desejada
	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}

//este código, a função fibonacci é definida de forma recursiva para calcular o número de Fibonacci na posição dada.
// O programa principal usa um loop para imprimir os números de Fibonacci até a posição desejada.
// Os comentários explicam cada parte do código
