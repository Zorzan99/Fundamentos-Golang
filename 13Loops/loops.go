package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second)

	}
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}
	nomes := [3]string{"joão", "gabriel", "zorzan"}

	//primeiro parametro sempre o indice, segundo sempre o parametro
	for indice, nome := range nomes {
		fmt.Println(indice, nome)

	}

	for indice1, letra := range "PALAVRA" {
		fmt.Println(indice1, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Joao",
		"sobrenome": "gabriel",
	}
	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
