package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Definindo a estrutura 'cachorro' com tags JSON para personalizar os nomes dos campos na serialização
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"chaveQualquer"`
}

func main() {
	// Criando uma instância da estrutura 'cachorro'
	c := cachorro{"Dominique", "Golden", 3}
	fmt.Println(c) // Imprimindo a estrutura original

	// Convertendo a estrutura 'cachorro' para formato JSON
	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	// Imprimindo a representação JSON da estrutura 'cachorro'
	fmt.Println(string(cachorroEmJson))

	// Imprimindo a representação JSON usando um buffer
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	// Criando um mapa representando um cachorro
	c2 := map[string]string{
		"nome": "Buddy",
		"raca": "Poodle",
	}

	// Convertendo o mapa para formato JSON
	cachorro2EmJson, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	// Imprimindo a representação JSON do mapa
	fmt.Println(string(cachorro2EmJson))

	// Imprimindo a representação JSON usando um buffer
	fmt.Println(bytes.NewBuffer(cachorro2EmJson))
}

// Definindo a Estrutura 'cachorro':

// Uma estrutura cachorro é definida com campos Nome, Raca e Idade. As tags json:"nome", json:"raca", e json:"chaveQualquer" são usadas para especificar os nomes dos campos na representação JSON.
// Criando uma Instância da Estrutura 'cachorro':

// Uma instância da estrutura cachorro chamada c é criada e impressa.
// Convertendo a Estrutura para JSON:

// A instância c é convertida para JSON usando json.Marshal, e a representação JSON é impressa.
// Usando um Buffer para a Representação JSON:

// A representação JSON de c é impressa usando um buffer.
// Criando um Mapa Representando um Cachorro:

// Um mapa c2 representando um cachorro é criado.
// Convertendo o Mapa para JSON:

// O mapa c2 é convertido para JSON usando json.Marshal, e a representação JSON é impressa.
// Usando um Buffer para a Representação JSON do Mapa:

// A representação JSON de c2 é impressa usando um buffer.
