package main

import "fmt"

func main() {
	// Operações aritméticas
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	var numero1 int16 = 10
	var numero2 int16 = 25

	// Conversão de tipos para realizar a soma
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)
	// FIM DAS OPERAÇÕES ARITMÉTICAS

	// Operadores de atribuição
	var variavel1 string = "String atribuição"
	variavel2 := "String 2 atribuicao"

	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// FIM DOS OPERADORES RELACIONAIS

	// Operadores lógicos
	fmt.Println("----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro, !falso)
	// FIM DOS OPERADORES LÓGICOS

	// Operadores unários
	numero := 10
	numero++
	numero += 15 // NUMERO = NUMERO + 15

	numero--
	numero -= 20 // NUMERO = NUMERO - 20

	numero *= 3 // NUMERO = NUMERO * 3

	numero /= 10
	numero %= 3

	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}

//Esse código apresenta operações aritméticas,
//operadores de atribuição, operadores relacionais,
//operadores lógicos e operadores unários.
