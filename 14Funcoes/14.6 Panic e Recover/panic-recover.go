package main

import "fmt"

// Função para recuperar a execução em caso de pânico (panic)
func recuperarExecucao() {
	// A função recover() é usada para recuperar a execução após um pânico
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar execução após pânico")
	}
}

// Função que verifica se um aluno está aprovado com base em duas notas
func alunoEstaAprovado(n1, n2 float64) bool {
	// O defer é usado para adiar a execução da função recuperarExecucao() até o final da função
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// Se a média for exatamente 6, é gerado um pânico
	panic("A média é exatamente 6")
}

func main() {
	// Chamada da função que verifica se um aluno está aprovado
	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("Pós execução")
}

//Neste código, a função recuperarExecucao é usada para tratar pânicos (panics) que podem ocorrer durante a execução do programa.
// A função alunoEstaAprovado calcula a média de duas notas e, caso a média seja exatamente 6, gera um pânico.
//O defer é usado para adiar a execução da função de recuperação (recuperarExecucao) até o final da função principal (main).
// Os comentários explicam cada parte do código.
