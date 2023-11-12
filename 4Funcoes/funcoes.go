package main

import "fmt"

// Função para somar dois números inteiros de 8 bits
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função para realizar operações matemáticas (soma e subtração) com dois números inteiros de 8 bits
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	// Chamando a função somar e imprimindo o resultado
	soma := somar(10, 20)
	fmt.Println(soma)

	// Declarando uma variável que armazena uma função anônima que imprime uma string
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	// Chamando a função anônima e imprimindo o resultado
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	// Chamando a função calculosMatematicos e imprimindo o resultado da soma
	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)
}
