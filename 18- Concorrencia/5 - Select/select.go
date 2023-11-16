package main

import (
	"fmt"
	"time"
)

func main() {
	// Criação de dois canais de strings
	fmt.Println("Select")
	canal1, canal2 := make(chan string), make(chan string)

	// Goroutine para enviar mensagens para canal1 a cada 500 milissegundos
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	// Goroutine para enviar mensagens para canal2 a cada 2 segundos
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// Loop principal
	for {
		// Select é usado para esperar múltiplos canais. O primeiro canal pronto será selecionado.
		select {
		// Se houver uma mensagem no canal1, imprime a mensagem
		case mensagemCanal1 := <-canal1:
			fmt.Printf("%s\n", mensagemCanal1)
		// Se houver uma mensagem no canal2, imprime a mensagem
		case mensagemCanal2 := <-canal2:
			fmt.Printf("%s\n", mensagemCanal2)
		}
	}
}

//A ideia geral é que existem duas goroutines em execução simultânea (criadas pelos go func()), cada uma enviando mensagens para um canal específico em intervalos regulares.
// O loop principal (for) está usando um select para aguardar mensagens de qualquer um dos canais.
// Quando uma mensagem é recebida em um dos canais, ela é impressa na tela.
// Como as goroutines estão executando em paralelo, você verá mensagens alternadas entre "Canal 1" e "Canal 2" na saída, dependendo dos intervalos de tempo especificados para cada goroutine.
