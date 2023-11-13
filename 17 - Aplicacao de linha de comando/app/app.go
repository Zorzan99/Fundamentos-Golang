package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerarvai retornar a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicacção de linha de comando"
	app.Usage = "Buscar IPS e Nomes de Servidos na internet."
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

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
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	//Pacote Net
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //NOME DO SERVER

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
