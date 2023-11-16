package main

import "fmt"

func main() {
	canal := make(chan string)
	//Canal com buffer só vai bloquear o trafego de dados quando o canal atingit sua capacidade máxima

	canal <- "Olá mundo"

	mensagem := <-canal
	fmt.Println(mensagem)

}
