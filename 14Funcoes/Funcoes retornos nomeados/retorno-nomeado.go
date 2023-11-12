package main

// Função calculosMatematicos realiza a soma e subtração de dois números inteiros
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	// Realiza a soma e a subtração dos números fornecidos
	soma = n1 + n2
	subtracao = n1 - n2

	// Retorna os resultados
	return
}

func main() {
	// Chamada da função calculosMatematicos com os valores 10 e 20
	// As variáveis soma e subtracao recebem os resultados
	soma, subtracao := calculosMatematicos(10, 20)

	// Imprime os resultados obtidos
	println("Soma:", soma, "Subtração:", subtracao)
}

//Neste código, a função calculosMatematicos recebe dois números inteiros, calcula a soma e a subtração e retorna esses valores.
// A função main chama a função calculosMatematicos com os valores 10 e 20, recebe os resultados nas variáveis soma e subtracao, e imprime esses resultados.
// O uso do := permite a inferência automática dos tipos das variáveis.
