package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando um canal para comunicação entre goroutines
	fmt.Println("Canais")
	canal := make(chan string)

	// Iniciando uma goroutine que executa a função escrever
	go escrever("Olá mundo!", canal)

	// Mensagem após o início da execução da goroutine escrever
	fmt.Println("Depois da função escrever começar a ser executada")

	// Loop infinito para receber mensagens do canal
	// for {
	// 	// Recebendo um valor do canal e verificando se o canal ainda está aberto
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break // Saindo do loop quando o canal estiver fechado
	// 	}
	// 	fmt.Println(mensagem)
	// }

	//range canal: Essa parte do loop utiliza a construção range para iterar sobre os valores enviados pelo canal.
	//Ele continuará iterando até que o canal seja fechado.

	//mensagem :=: Isso declara e inicializa a variável mensagem com o valor recebido do canal em cada iteração.

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	// Mensagem de encerramento do programa
	fmt.Println("Fim do programa")
}

// Função que escreve uma mensagem no canal repetidamente
func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Enviando um valor para o canal
		time.Sleep(time.Second)
	}
	close(canal) // Fechando o canal após o envio das mensagens
}

//Neste código em Go, é criado um canal (canal) para a comunicação entre duas goroutines.
// A função main inicia uma goroutine chamando a função escrever, que envia a mensagem "Olá mundo!" para o canal a cada segundo, cinco vezes.
// Enquanto isso, a goroutine principal (main) aguarda e imprime as mensagens recebidas do canal até que o canal seja fechado pela goroutine escrever.
// O fechamento do canal é feito através da função close(canal) após o envio das cinco mensagens.
// O programa então imprime "Fim do programa".
//A expressão mensagem, aberto := <-canal tenta receber um valor do canal. Se o canal estiver aberto e houver uma mensagem para ser lida, a variável mensagem receberá o valor do canal, e aberto será true.
// Se o canal estiver fechado e não houver mais mensagens para ler, aberto será false, e o loop será encerrado.
//Portanto, aberto indica se o canal está aberto ou fechado, ajudando a controlar a execução do loop que lê mensagens do canal na função main.
