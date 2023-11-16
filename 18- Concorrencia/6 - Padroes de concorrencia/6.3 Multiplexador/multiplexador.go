package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Função principal
func main() {
	fmt.Println("Multiplexador")

	// Chama a função multiplexador para combinar mensagens de dois canais
	canal := multiplexador(escrever("Olá mundo"), escrever("Programango em Go"))

	// Recebe e imprime mensagens do canal combinado 5 vezes
	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

// Função multiplexador
func multiplexador(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	// Cria um canal de saída
	canalDeSaida := make(chan string)

	// Inicia uma goroutine para combinar mensagens dos canais de entrada
	go func() {
		for {
			// Seleciona a mensagem de qualquer canal disponível usando select
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	// Retorna o canal de saída
	return canalDeSaida
}

// Função escrever
func escrever(texto string) <-chan string {
	// Cria um canal de string
	canal := make(chan string)

	// Inicia uma goroutine para enviar mensagens ao canal em um loop
	go func() {
		for {
			// Formata a mensagem e envia para o canal
			canal <- fmt.Sprintf("valor recebido: %s", texto)

			// Aguarda um tempo aleatório antes de enviar a próxima mensagem
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(2000)))
		}
	}()

	// Retorna o canal
	return canal
}

// Explicação:
// main:
// Chama a função multiplexador para combinar mensagens de dois canais.
// Entra em um loop para receber e imprimir mensagens do canal combinado 5 vezes.
// multiplexador:
// Cria um canal de saída.
// Inicia uma goroutine que usa a instrução select para selecionar mensagens de dois canais de entrada diferentes e enviá-las para o canal de saída.
// escrever:
// Cria um canal de string.
// Inicia uma goroutine que envia mensagens formatadas ao canal em um loop, com pausas aleatórias entre as mensagens.
// Este programa simula a combinação de mensagens de dois canais usando um multiplexador, e a função main imprime algumas dessas mensagens combinadas.
// O uso de pausas aleatórias na função escrever adiciona um componente de assincronia ao processo.
