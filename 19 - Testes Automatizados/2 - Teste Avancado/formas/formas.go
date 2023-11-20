package formas

import (
	"math"
)

// Definição da interface "forma" com um método "area()"
type Forma interface {
	area() float64
}

// Função que escreve a área da forma usando a interface "forma"

// Definição da estrutura "retangulo"
type Retangulo struct {
	altura  float64
	largura float64
}

// Método "area()" associado à estrutura "retangulo"
func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

// Definição da estrutura "circulo"
type Circulo struct {
	raio float64
}

// Método "area()" associado à estrutura "circulo"
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

//Neste código, é definida uma interface forma com um método area().
// As estruturas retangulo e circulo implementam essa interface ao definir seus próprios métodos area().
// A função escreverArea aceita qualquer tipo que implemente a interface forma e imprime a área da forma passada.
// Os comentários explicam cada parte do código.
