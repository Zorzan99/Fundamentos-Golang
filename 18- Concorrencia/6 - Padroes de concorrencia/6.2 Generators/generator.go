package main

import (
	"fmt"
	"time"
)

// Função principal
func main() {
	fmt.Println("Generator")

	// Chama a função escrever para obter um canal
	canal := escrever("OLÁ MUNDO")

	// Recebe e imprime mensagens do canal 10 vezes
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Função escrever
func escrever(texto string) <-chan string {
	// Cria um canal de string
	canal := make(chan string)

	// Inicia uma goroutine para enviar mensagens ao canal
	go func() {
		for {
			// Formata a mensagem e envia para o canal
			canal <- fmt.Sprintf("valor recebido:  %s", texto)

			// Aguarda por 500 milissegundos antes de enviar a próxima mensagem
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// Retorna o canal
	return canal
}

// Explicação:

// main:

// Chama a função escrever para obter um canal.
// Entra em um loop para receber e imprimir mensagens do canal 10 vezes.
// escrever:
// Cria um canal de string.
// Inicia uma goroutine que envia mensagens formatadas ao canal em um loop infinito.
// Aguarda 500 milissegundos entre cada envio.
// Retorna o canal.
// Esse programa cria um gerador de mensagens assíncrono.
// A função main inicia a produção de mensagens chamando a função escrever e, em seguida, imprime 10 dessas mensagens à medida que são recebidas do canal.
// Como a goroutine continua executando em paralelo, as mensagens são geradas e consumidas concorrentemente.
