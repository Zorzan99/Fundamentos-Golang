package main

import "fmt"

// Função que imprime uma mensagem indicando que está executando a função1
func funcao1() {
	fmt.Println("Executando a função1")
}

// Função que imprime uma mensagem indicando que está executando a função2
func funcao2() {
	fmt.Println("Executando a função2")
}

// Função que verifica se um aluno está aprovado com base em duas notas
func alunoEstaAprovado(n1, n2 float32) bool {
	// O defer é usado para adiar a execução da mensagem para depois do retorno da função
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	// Cálculo da média
	media := (n1 + n2) / 2

	// Verificação se a média é maior ou igual a 6
	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Defer")

	// O defer é usado para adiar a execução da função1 até o final do escopo da função main
	defer funcao1()
	// A função2 é executada antes da função1, mas o defer faz com que a função1 seja adiada para o final
	funcao2()

	// Chamada da função que verifica se um aluno está aprovado
	fmt.Println(alunoEstaAprovado(8, 8))
}

//Neste código, são utilizadas funções para imprimir mensagens e verificar se um aluno está aprovado com base em suas notas.
// O defer é usado para adiar a execução de uma função até o final do escopo da função que a contém.
//Os comentários explicam cada parte do código
