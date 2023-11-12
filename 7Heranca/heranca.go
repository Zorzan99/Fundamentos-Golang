package main

import "fmt"

// Definição da estrutura "pessoa"
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// Definição da estrutura "estudante" que incorpora a estrutura "pessoa"
type estudante struct {
	pessoa    // Incorporação da estrutura "pessoa" na estrutura "estudante"
	curso     string
	faculdade string
}

func main() {
	// Impressão de texto informativo
	fmt.Println("Herança")

	// Criação de uma instância da estrutura "pessoa"
	p1 := pessoa{"João", "Gabriel", 20, 173}
	fmt.Println(p1)

	// Criação de uma instância da estrutura "estudante" com valores específicos
	e1 := estudante{p1, "ADS", "unip"}
	fmt.Println(e1.nome) // Acesso ao campo "nome" da estrutura incorporada "pessoa"
}
