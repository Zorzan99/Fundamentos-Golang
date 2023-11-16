package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Cria uma WaitGroup para aguardar a conclusão das goroutines
	var waitGroup sync.WaitGroup

	// Adiciona dois contadores à WaitGroup, indicando que esperamos a conclusão de duas goroutines. Sempre especificar a quantidade de Go Routines pois ele nao reconhece sozinho
	waitGroup.Add(2)

	// Inicia uma goroutine para chamar a função escrever com o texto "Olá mundo"
	go func() {
		escrever("Olá mundo")
		// Quando a goroutine termina, sinaliza à WaitGroup que ela concluiu (-1)
		waitGroup.Done()
	}()

	// Inicia outra goroutine para chamar a função escrever com o texto "Programando em Go!"
	go func() {
		escrever("Programando em Go!")
		// Quando a goroutine termina, sinaliza à WaitGroup que ela concluiu (-1)
		waitGroup.Done()
	}()

	// Aguarda até que a WaitGroup atinja 0, o que significa que ambas as goroutines foram concluídas
	waitGroup.Wait()
}

// escrever é uma função que imprime repetidamente um texto com um intervalo de 1 segundo
func escrever(texto string) {
	// O loop externo executa 5 vezes
	for i := 0; i < 5; i++ {
		// O loop interno imprime o texto recebido como parâmetro
		fmt.Println(texto)
		// Dorme por 1 segundo antes de continuar para criar um intervalo entre as impressões
		time.Sleep(time.Second)
	}
}

// 1 - main:
// Cria uma sync.WaitGroup para coordenar a execução de goroutines.
// Adiciona dois contadores à WaitGroup usando waitGroup.Add(2), indicando que esperamos a conclusão de duas goroutines.
// Inicia duas goroutines anônimas usando go func() { ... }, cada uma chamando a função escrever com um texto específico.
// Cada goroutine chama waitGroup.Done() quando termina, sinalizando à WaitGroup que ela concluiu (-1).
// waitGroup.Wait() aguarda até que a WaitGroup atinja 0, o que significa que ambas as goroutines foram concluídas.
// 2 - escrever:
// A função escrever é semelhante à versão anterior, mas agora tem um loop externo que executa 5 vezes para simular uma operação mais longa.
// Isso demonstra como a sync.WaitGroup pode ser usada para aguardar a conclusão de múltiplas goroutines antes de prosseguir.
// Essencialmente, a sync.WaitGroup é usada para coordenar a execução concorrente e aguardar que todas as goroutines concluam antes de continuar.
