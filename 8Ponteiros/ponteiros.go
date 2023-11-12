package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	// Declaração e inicialização de variáveis
	var variavel1 int = 10
	var variavel2 int = variavel1

	// Impressão dos valores iniciais das variáveis
	fmt.Println(variavel1, variavel2)

	// Incremento da variável1
	variavel1++

	// Impressão dos valores após o incremento
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int // Guarda o valor inteiro
	var ponteiro *int // Guarda o endereço de memória de um valor inteiro

	variavel3 = 100
	ponteiro = &variavel3 // Atribuição do endereço de memória da variável3 ao ponteiro

	// Impressão do valor e do ponteiro
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // Processo de desreferenciação, desfazendo a referência

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // O valor associado ao ponteiro também é modificado

}

//Este código demonstra o conceito de ponteiros em Go.
//Os comentários explicam as operações realizadas, incluindo a criação de variáveis, a atribuição de valores e endereços de memória, e o processo de desreferenciação (usando o operador * para acessar o valor associado ao ponteiro).
