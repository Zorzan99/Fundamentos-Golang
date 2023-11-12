package main

import "fmt"

// Definição da estrutura "endereco"
type endereco struct {
	logradouro string
	numero     uint8
}

// Definição da estrutura "usuario" que incorpora a estrutura "endereco"
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	// Impressão de texto informativo
	fmt.Println("Arquivo structs")

	// Declaração de uma variável do tipo "usuario"
	var u usuario

	// Atribuição de valores aos campos da estrutura "usuario"
	u.nome = "Gabriel"
	u.idade = 24

	// Impressão da estrutura "usuario"
	fmt.Println(u)

	// Impressão de campos específicos da estrutura "usuario"
	fmt.Println("Nome", u.nome, "Idade", u.idade)

	// Criação de uma instância da estrutura "endereco" e atribuição a um campo da estrutura "usuario"
	enderecoExemplo := endereco{"Rua teste", 0}

	// Criação de uma instância da estrutura "usuario" com valores específicos
	usuario2 := usuario{"Gabriel", 24, enderecoExemplo}
	fmt.Println(usuario2)

	// Criação de uma instância da estrutura "usuario" com valor específico para o campo "idade"
	usuario3 := usuario{idade: 24}
	fmt.Println(usuario3)
}

//Esse código trabalha com a definição e utilização de estruturas em Go.
//Ele define duas estruturas (endereco e usuario) e demonstra como criar e manipular instâncias dessas estruturas,
//bem como imprimir informações relevantes sobre elas
