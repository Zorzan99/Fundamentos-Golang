package main

import (
	"errors"
	"fmt"
)

func main() {
	// Declaração e inicialização de inteiros
	var numero int64 = 1000000000000000000
	fmt.Println(numero)

	var numero1 int = 1000000000
	fmt.Println(numero1)

	numero2 := 10000000000000
	fmt.Println(numero2)

	// ALIAS: INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Declaração e inicialização de números reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.4555
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// FIM NUMEROS REAIS

	// Declaração e inicialização de strings
	var str string = "Texto str"
	fmt.Println(str)

	str2 := "Texto2 str"
	fmt.Println(str2)

	// FIM STRINGS

	// Declaração e inicialização de booleanos
	var booleano1 bool = true
	fmt.Println(booleano1)

	// para inicializar como false
	var booleano2 bool
	fmt.Println(booleano2)

	// FIM BOOLEANOS

	// Valor padrão é NIL (zero value) e não NEW
	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}

// 1 O código começa com package main, indicando que este é um programa executável e deve incluir uma função main.

// 2 import "fmt" e import "errors": Importa os pacotes fmt para formatação de saída e errors para manipulação de erros.

// 3 Diversos exemplos de declaração e inicialização de variáveis dos tipos:

// 4 int64, int, rune, byte: Números inteiros com diferentes tamanhos e aliases.
// 5 float32, float64: Números reais com diferentes precisões.
// 6 string: Sequências de caracteres.
// 7 bool: Valores booleanos (true/false).
// 8 error: Tipo especial para representar erros.
// 9 Utilização de aliases como rune (um alias para int32) e byte (um alias para uint8).

// 10 Demonstração de como inicializar uma variável booleana com valor padrão (false).

// 11 Uso do tipo error para representar erros. A variável erro é inicializada com o valor padrão (nil), enquanto erro1 é inicializada com um erro personalizado usando errors.New.
