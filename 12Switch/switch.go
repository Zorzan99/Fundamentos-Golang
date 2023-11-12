package main

import "fmt"

// Função que retorna o dia da semana com base no número passado como parâmetro
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

// Função equivalente à anterior, mas usando switch sem expressão
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

// Função equivalente às anteriores, mas usando variável intermediária para armazenar o dia da semana
func diaDaSemana3(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		return "Número inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	// Chamadas das funções e impressão dos resultados
	dia := diaDaSemana(7)
	fmt.Println(dia)

	dia2 := diaDaSemana2(6)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(10)
	fmt.Println(dia3)
}

//Neste código, são utilizadas três abordagens diferentes para implementar uma função que retorna o dia da semana com base em um número passado como parâmetro.
// As funções utilizam a estrutura de controle switch de maneiras ligeiramente diferentes, e os comentários explicam cada abordagem.
