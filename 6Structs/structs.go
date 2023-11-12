package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")
	var u usuario

	u.nome = "Gabriel"
	u.idade = 24

	fmt.Println(u)
	fmt.Println("Nome", u.nome, "Idade", u.idade)

	enderecoExemplo := endereco{"Rua teste", 0}

	usuario2 := usuario{"Gabriel", 24, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 24}
	fmt.Println(usuario3)

}
