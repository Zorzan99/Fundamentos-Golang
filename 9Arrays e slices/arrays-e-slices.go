package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// DECLARAÇÃO E INICIALIZAÇÃO DE ARRAYS
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "Gabriel"
	fmt.Println(array2)

	array3 := [5]string{"Posição1", "Posição2", "Posição3", "Posição4", "Posição5"}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)

	// DECLARAÇÃO E INICIALIZAÇÃO DE SLICES
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// OBTENDO O TIPO DE UMA VARIÁVEL COM REFLECT
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array4))

	// ADICIONANDO ELEMENTO A UM SLICE COM APPEND
	slice = append(slice, 20)
	fmt.Println(slice)

	// CRIAÇÃO DE UM SLICE A PARTIR DE UM ARRAY
	slice2 := array3[1:3]
	fmt.Println(slice2)

	// MODIFICANDO ELEMENTO DO ARRAY ORIGINAL AFETA O SLICE
	array3[1] = "Posição Alterada"
	fmt.Println(slice2)

	// ARRAYS INTERNOS (SLICES COM MAKE)
	fmt.Println("-----------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // comprimento (length)
	fmt.Println(cap(slice3)) // capacidade

	// ADICIONANDO ELEMENTO A UM SLICE COM APPEND (AUTOMATICAMENTE REALOCAÇÃO)
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}

//Neste código, são abordados conceitos relacionados a arrays e slices em Go.
//Os comentários explicam a declaração, inicialização, manipulação e características de arrays e slices.
// Além disso, são mostradas as diferenças entre arrays e slices, bem como o uso do make para criar slices com capacidade específica.
