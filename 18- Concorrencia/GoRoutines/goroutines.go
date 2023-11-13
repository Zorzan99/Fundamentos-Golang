package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO

	// Inicia uma goroutine para chamar a função escrever com o texto "Olá mundo"
	go escrever("Olá mundo") // goroutine

	// Chama a função escrever diretamente com o texto "Programando em Go!"
	escrever("Programando em Go!")
}

// escrever é uma função que imprime repetidamente um texto com um intervalo de 1 segundo
func escrever(texto string) {
	// O loop infinito executa continuamente
	for {
		// Imprime o texto recebido como parâmetro
		fmt.Println(texto)
		// Dorme por 1 segundo antes de repetir o loop
		time.Sleep(time.Second)
	}
}
