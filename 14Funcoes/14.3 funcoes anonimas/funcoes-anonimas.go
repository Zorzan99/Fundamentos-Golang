package main

import "fmt"

func main() {
	fmt.Println("Funções anonimas")

	func() {
		fmt.Println("Olá mundo")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("PASSANDO PARAMETRO")

	retorno := func(texto string) string {
		return fmt.Sprint("Recebido -> ", texto)
	}("PASSANDO PARAMETRO2")

	fmt.Println(retorno)

}
