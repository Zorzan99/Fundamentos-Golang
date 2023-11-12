package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// MAP SIMPLES
	usuario := map[string]string{
		"nome":      "Gabriel",
		"sobrenome": "Zorzan",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	// MAPS ANINHADOS
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

	// DELETANDO UMA CHAVE DE UM MAP
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}

//Neste código, são demonstrados conceitos relacionados a maps (mapas) em Go.
//Os comentários explicam a criação e manipulação de mapas, incluindo mapas aninhados.
//Além disso, é mostrado como deletar uma chave de um mapa usando a função delete.
