package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	// ESTRUTURA DE CONTROLE IF-ELSE
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// ESTRUTURA DE CONTROLE IF-ELSE COM INICIALIZAÇÃO DE VARIÁVEL
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}
}

//Este código demonstra o uso de estruturas de controle if-else em Go.
// Os comentários explicam cada parte do código, incluindo a sintaxe do if-else simples e a versão com inicialização de variável antes da condiçã
