package main

import (
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
	// Representação JSON de um cachorro
	cachorroEmJson := `{"nome":"Buddy","raca":"Poodle"}`

	// Declarando uma variável do tipo cachorro
	var c cachorro

	// Decodificando a representação JSON em uma estrutura 'cachorro'
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	// Imprimindo a estrutura 'cachorro' após a decodificação
	fmt.Println(c)

	// Representação JSON de um mapa representando um cachorro
	cachorro2EmJson := `{"nome": "Dominique", "raca":"Golden"}`

	// Declarando um mapa
	c2 := make(map[string]string)

	// Decodificando a representação JSON em um mapa
	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	// Imprimindo o mapa após a decodificação
	fmt.Println(c2)
}

// Definindo a Estrutura 'cachorro':

// É definida uma estrutura cachorro com campos Nome, Raca, e Idade. As tags json:"nome", json:"raca", e json:"chaveQualquer" são usadas para personalizar os nomes dos campos durante a desserialização.
// Decodificando JSON para a Estrutura 'cachorro':

// Uma representação JSON de um cachorro é declarada como uma string (cachorroEmJson).
// A função json.Unmarshal é usada para desserializar a string JSON em uma instância da estrutura cachorro.
// Imprimindo a Estrutura 'cachorro' Decodificada:

// A estrutura c é impressa após a desserialização.
// Decodificando JSON para um Mapa:

// Uma representação JSON de um cachorro (agora representado como um mapa) é declarada como uma string (cachorro2EmJson).
// A função json.Unmarshal é usada para desserializar a string JSON em um mapa.
// Imprimindo o Mapa Decodificado:

// O mapa c2 é impresso após a desserialização.
// Este programa ilustra como usar a função json.Unmarshal para desserializar dados JSON em estruturas de dados Go.
