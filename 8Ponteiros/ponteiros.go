package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERENCIA DE MEMÓRIA
	var variavel3 int //guarda o valor inteiro
	var ponteiro *int //guarda o endereco de memoria de um valor inteiro

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //Processo de desreferenciação, desfazendo a referencia

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)

}
