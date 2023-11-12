package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	// LOOP COM ESTRUTURA WHILE
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second)
	}

	// LOOP COM ESTRUTURA FOR CLÁSSICO
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	// LOOP COM RANGE EM UM ARRAY
	nomes := [3]string{"joão", "gabriel", "zorzan"}
	// primeiro parâmetro sempre o índice, segundo sempre o elemento
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// LOOP COM RANGE EM UMA STRING
	for indice1, letra := range "PALAVRA" {
		fmt.Println(indice1, string(letra))
	}

	// LOOP COM RANGE EM UM MAP
	usuario := map[string]string{
		"nome":      "Joao",
		"sobrenome": "gabriel",
	}
	fmt.Println(usuario)

	// primeiro parâmetro sempre a chave, segundo sempre o valor
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}

// Neste código, são apresentados diferentes tipos de loops em Go, incluindo loops usando estruturas for, for com a condição range em arrays, strings e maps.
//Os comentários explicam cada parte do código.
