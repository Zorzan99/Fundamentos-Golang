package main

import "fmt"

// Função principal
func main() {
	fmt.Println("Worker-Pools")

	// Canais para enviar tarefas e receber resultados
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Inicia quatro goroutines worker
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	// Envia tarefas para o canal
	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas) // Fecha o canal de tarefas para indicar que não há mais tarefas a serem enviadas

	// Recebe e imprime os resultados
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// Função worker
func worker(tarefas <-chan int, resultados chan<- int) {
	// Loop para receber tarefas do canal e enviar resultados
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

// Função para calcular o número de Fibonacci
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	// Chama recursivamente a função para as posições anteriores e soma os resultados
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//Explicação:

// main:

// Cria canais (tarefas e resultados) para coordenar a comunicação entre a função main e as goroutines worker.
// Inicia quatro goroutines worker usando a função go worker(tarefas, resultados).
// Envia tarefas (números) para o canal tarefas.
// Fecha o canal tarefas para indicar que não há mais tarefas a serem enviadas.
// Recebe e imprime os resultados do canal resultados.
// worker:
// Recebe tarefas do canal tarefas e envia os resultados para o canal resultados.
// fibonacci:
// Calcula o número de Fibonacci para uma posição dada usando recursão.
// O programa cria um Worker Pool para calcular os números de Fibonacci de maneira concorrente, melhorando a eficiência do processamento.
// Cada worker recebe tarefas do canal, realiza o cálculo e envia os resultados de volta para o canal.
