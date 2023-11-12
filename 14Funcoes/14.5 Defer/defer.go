package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função1")
}

func funcao2() {
	fmt.Println("Executando a função2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado sera retornado")
	fmt.Println("Entrando na função para verificar se o aluno esta aprovvado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false

}

func main() {
	fmt.Println("Defer")
	//defer funcao1()
	// //DEFER = ADIAR
	// funcao2()
	fmt.Print(alunoEstaAprovado(8, 8))
}
