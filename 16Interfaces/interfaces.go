package main

import (
	"fmt"
	"math"
)

// Definição da interface "forma" com um método "area()"
type forma interface {
	area() float64
}

// Função que escreve a área da forma usando a interface "forma"
func escreverArea(f forma) {
	fmt.Println("A área da forma é", f.area())
}

// Definição da estrutura "retangulo"
type retangulo struct {
	altura  float64
	largura float64
}

// Método "area()" associado à estrutura "retangulo"
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// Definição da estrutura "circulo"
type circulo struct {
	raio float64
}

// Método "area()" associado à estrutura "circulo"
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	fmt.Println("Interfaces")

	// Criação de uma instância da estrutura "retangulo"
	r := retangulo{10, 15}

	// Chamada da função escreverArea passando a instância da estrutura "retangulo"
	escreverArea(r)

	// Criação de uma instância da estrutura "circulo"
	c := circulo{10}

	// Chamada da função escreverArea passando a instância da estrutura "circulo"
	escreverArea(c)
}

//Neste código, é definida uma interface forma com um método area().
// As estruturas retangulo e circulo implementam essa interface ao definir seus próprios métodos area().
// A função escreverArea aceita qualquer tipo que implemente a interface forma e imprime a área da forma passada.
// Os comentários explicam cada parte do código.
