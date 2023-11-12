package main

import "fmt"

// Definição da estrutura de dados "usuario"
type usuario struct {
	nome  string
	idade uint8
}

// Método salvar associado à estrutura "usuario"
func (u usuario) salvar() {
	fmt.Println("Salvando o Usuário:", u.nome, "de idade", u.idade)
}

// Método verificarIdade associado à estrutura "usuario"
func (u usuario) verificarIdade() string {
	if u.idade >= 18 {
		return "Maior de idade"
	} else {
		return "Menor de idade"
	}
}

// Método fazerAniversario associado à estrutura "usuario" com ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")

	// Criação de uma instância da estrutura "usuario"
	usuario1 := usuario{"Usuário1", 7}
	fmt.Println(usuario1)

	// Chamada do método salvar associado à instância da estrutura "usuario"
	usuario1.salvar()

	// Chamada do método verificarIdade associado à instância da estrutura "usuario"
	maiorDeIdade := usuario1.verificarIdade()
	fmt.Println(maiorDeIdade)

	// Chamada do método fazerAniversario associado à instância da estrutura "usuario" (modifica o valor com ponteiro)
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}

//Neste código, é definida uma estrutura de dados usuario com métodos associados a ela.
// Os métodos (salvar, verificarIdade, e fazerAniversario) são associados à estrutura usuario e podem ser chamados usando uma instância dessa estrutura.
// O método fazerAniversario é definido com um ponteiro para a estrutura para que possa modificar o valor da idade da instância original.
// Os comentários explicam cada parte do código.
