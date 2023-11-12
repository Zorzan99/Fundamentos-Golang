package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int64 = 1000000000000000000
	fmt.Println(numero)

	var numero1 int = 1000000000
	fmt.Println(numero1)

	numero2 := 10000000000000
	fmt.Println(numero2)

	//ALIAS
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.4555
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//FIM NUMEROS REAIS

	var str string = "Texto str"
	fmt.Println(str)

	str2 := "Texto2 str"
	fmt.Println(str2)

	//FIM STRINGS

	var booleano1 bool = true
	fmt.Println(booleano1)

	//para inicializar como false
	var booleano2 bool
	fmt.Println(booleano2)

	// FIM BOOLEANOS

	//valor padrao nele é NIL de 0 e nao NEW de novo
	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
