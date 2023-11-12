package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando o Usuario: ", u.nome, " de idade ", u.idade)
}

func (u usuario) verificarIdade() string {
	if u.idade >= 18 {
		return "Maior de idade"
	} else {
		return "Menor de idade"
	}
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")
	usuario1 := usuario{"Usuário1", 7}
	fmt.Println(usuario1)
	usuario1.salvar()

	maiorDeIdade := usuario1.verificarIdade()
	fmt.Println(maiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
