package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	// Cria uma nova instância de aplicação
	app := cli.NewApp()

	// Configura o nome e a descrição da aplicação
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Buscar IPS e Nomes de Servidores na internet."

	// Define uma flag para o host, com um valor padrão
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	// Define os comandos disponíveis na aplicação
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca nomes dos servidores da internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	// Retorna a instância da aplicação, que está pronta para ser executada
	return app
}

// buscarIps é a função associada ao comando "ip"
func buscarIps(c *cli.Context) {
	// Obtém o valor da flag "host" do contexto da linha de comando
	host := c.String("host")

	// Realiza uma busca de IPs associados ao host fornecido
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	// Imprime os IPs encontrados
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// buscarServidores é a função associada ao comando "servidores"
func buscarServidores(c *cli.Context) {
	// Obtém o valor da flag "host" do contexto da linha de comando
	host := c.String("host")

	// Realiza uma busca de nomes de servidores associados ao host fornecido
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	// Imprime os nomes dos servidores encontrados
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

//Função Gerar:
// Gerar é uma função que cria e retorna uma instância de cli.App, que representa a aplicação de linha de comando.
// cli.NewApp() cria uma nova instância de aplicação.
// app.Name define o nome da aplicação ("Aplicação de linha de comando").
// app.Usage define uma breve descrição do propósito da aplicação ("Buscar IPS e Nomes de Servidores na internet.").
// flags é uma lista de flags, neste caso, contendo apenas uma flag host.
// app.Commands define os comandos disponíveis na aplicação. No caso, há dois comandos: ip e servidores, ambos usando a flag host e associados às funções buscarIps e buscarServidores, respectivamente.
// A função retorna a instância da aplicação, que está pronta para ser executada.
// Função buscarIps:

// buscarIps é a função associada ao comando ip.
// c.String("host") obtém o valor da flag host do contexto da linha de comando.
// net.LookupIP(host) realiza uma busca de IPs associados ao host fornecido.
// Os IPs encontrados são impressos no console.
// Função buscarServidores:

// buscarServidores é a função associada ao comando servidores.
// c.String("host") obtém o valor da flag host do contexto da linha de comando.
// net.LookupNS(host) realiza uma busca de nomes de servidores associados ao host fornecido.
// Os nomes dos servidores encontrados são impressos no console.
// O aplicativo, quando executado, permite que o usuário escolha entre os comandos ip e servidores, fornecendo um host opcional através da flag --host. O código faz uso da biblioteca urfave/cli para facilitar a criação de aplicativos de linha de comando em Go.
