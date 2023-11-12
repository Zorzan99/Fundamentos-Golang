package main

import "fmt"

// Função que retorna uma função (closure)
func closure() func() {
	// Variável local à função closure
	texto := "Dentro da função closure"

	// Função anônima (closure) que imprime o texto da variável acima
	funcao := func() {
		fmt.Println(texto)
	}

	// Retorna a função anônima (closure)
	return funcao
}

func main() {
	// Variável local à função main
	texto := "Dentro da função main"
	fmt.Println(texto)

	// Chama a função closure(), que retorna a função anônima (closure)
	funcaoNova := closure()

	// Chama a função anônima (closure) armazenada na variável funcaoNova
	funcaoNova()
}

//Neste código, é demonstrado o conceito de closure em Go.
// A função closure retorna uma função anônima que imprime o valor da variável texto, que é uma variável local dentro da função closure.
// A função main também possui uma variável texto, mas a função anônima ainda tem acesso à variável texto da função closure porque ela é um closure.
// Os comentários explicam cada parte do códig
