package main

import "fmt"

// Função que inverte o sinal de um número inteiro e retorna o resultado
func inverterSinal(numero int) int {
	return numero * -1
}

// Função que inverte o sinal de um número inteiro usando ponteiros
func inverterSinalComPonteiro(numero *int) {
	// Desreferenciação do ponteiro para modificar o valor do número diretamente na memória
	*numero = *numero * -1
}

func main() {
	fmt.Println("Ponteiros")

	// Número inicial
	numero := 20

	// Chamada da função inverterSinal para inverter o sinal do número
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido) // Resultado da função (não modifica o número original)
	fmt.Println(numero)          // Número original (não é modificado)

	// Novo número
	novoNumero := 40
	fmt.Println(novoNumero) // Número original antes de chamar a função

	// Chamada da função inverterSinalComPonteiro para inverter o sinal do número usando ponteiro
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero) // Número modificado pela função usando ponteiro
}

//Neste código, é demonstrado o uso de ponteiros para modificar o valor de uma variável diretamente na memória.
// A função inverterSinal retorna um novo valor com o sinal invertido, enquanto a função inverterSinalComPonteiro modifica o valor original usando um ponteiro.
// Os comentários explicam cada parte do código
