package main

import "fmt"

// Função genérica que aceita qualquer tipo de parâmetro através de uma interface vazia
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Interfaces tipo genérico")

	// Chamadas da função genérica com diferentes tipos de parâmetros
	generica("String") // Uma string
	generica(1)        // Um número inteiro
	generica(true)     // Um valor booleano
	generica(87.3232)  // Um número de ponto flutuante
}

//Neste código, a função generica é uma função genérica que aceita qualquer tipo de parâmetro através de uma interface vazia (interface{}).
// Isso permite que ela seja usada com diferentes tipos de dados, como strings, inteiros, booleanos e números de ponto flutuante.
// Os comentários explicam cada parte do código.
// Essa técnica é conhecida como "empty interface" ou "interface{}" e é comumente usada em Go para criar funções ou estruturas de dados mais flexíveis em termos de tipos de dados aceitos.
