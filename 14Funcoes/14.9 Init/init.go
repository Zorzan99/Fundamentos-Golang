package main

import "fmt"

// Variável global n
var n int

// Função init é executada automaticamente antes da função main
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n) // A variável n já foi inicializada pela função init
}

//Neste código, a função init é executada automaticamente antes da função main quando o programa é iniciado.
// A função init inicializa a variável global n com o valor 10. Em seguida, a função main é executada, imprimindo a mensagem "Função main sendo executada" e o valor da variável n.
// Os comentários explicam cada parte do código.
