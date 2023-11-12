package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

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

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array4))

	slice = append(slice, 20)
	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)
	array3[1] = "Posição Alterada"
	fmt.Println(slice2)
	////

	///ARRAYS INTERNOS

	fmt.Println("-----------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade

	//array uma lista com tamanho fixo, slice uma array sem tamnho fixo
}
