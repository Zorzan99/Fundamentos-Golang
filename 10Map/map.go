package main

import "fmt"

func main() {
	fmt.Println("Maps")
	usuario := map[string]string{
		"nome":      "Gabriel",
		"sobrenome": "Zorzan",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiroNome": "Gabriel",
			"segundoNome":  "Zorzan",
		},
		"curso": {
			"nomeCurso": "Engenharia de Software",
			"campus":    "Campus1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
