package main

func calculosMatematicos(n1, n2 int) (soma int, subtratacao int) {
	soma = n1 + n2
	subtratacao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	println(soma, subtracao)
}
