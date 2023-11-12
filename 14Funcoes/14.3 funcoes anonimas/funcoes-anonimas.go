package main

import "fmt"

func main() {
	fmt.Println("Funções anônimas")

	// Função anônima sem parâmetros
	func() {
		fmt.Println("Olá mundo")
	}() // A chamada dos parênteses () executa a função imediatamente

	// Função anônima com parâmetro
	func(texto string) {
		fmt.Println(texto)
	}("PASSANDO PARAMETRO")

	// Função anônima com retorno
	retorno := func(texto string) string {
		return fmt.Sprint("Recebido -> ", texto)
	}("PASSANDO PARAMETRO2")

	fmt.Println(retorno)
}

//Neste código, são utilizadas funções anônimas em Go.
// Os comentários explicam cada parte do código, incluindo exemplos de funções anônimas sem parâmetros, com parâmetros e com retorno.
//A chamada das funções anônimas é realizada imediatamente após a sua definição.
