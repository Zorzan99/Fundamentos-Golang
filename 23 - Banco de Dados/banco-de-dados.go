// O pacote "main" é o ponto de entrada principal para um programa em Go.
package main

// Importação de pacotes necessários.
// "database/sql" é usado para interagir com bancos de dados SQL.
// "fmt" fornece funcionalidades de formatação e impressão.
// "log" é utilizado para log de mensagens.
// O pacote "_" é usado para importar um pacote sem usá-lo explicitamente no código,
// nesse caso, o driver MySQL "github.com/go-sql-driver/mysql".
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Função principal chamada ao executar o programa.
func main() {
	// String de conexão com o banco de dados MySQL.
	urlConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	// Abertura de uma conexão com o banco de dados MySQL usando o driver "mysql".
	// O método sql.Open retorna um objeto de banco de dados e um possível erro.
	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		// Em caso de erro ao abrir a conexão, o programa loga a mensagem de erro e encerra.
		log.Fatal(erro)
	}

	// O "defer" garante que a função "db.Close()" será chamada antes do término da função "main".
	defer db.Close()

	// Verifica se é possível pingar o banco de dados para garantir que a conexão está ativa.
	// Se houver algum problema, o programa loga a mensagem de erro e encerra.
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	// Se tudo estiver correto até este ponto, imprime uma mensagem indicando que a conexão está aberta.
	fmt.Println("Conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}

// Explicação do Código:

// Importações:

// database/sql: Pacote para interagir com bancos de dados SQL.
// fmt: Pacote para formatação e impressão.
// log: Pacote para registrar mensagens de log.
// _ "github.com/go-sql-driver/mysql": Importação do driver MySQL sem usar explicitamente no código (o "_" é usado para evitar um erro de compilação por não usar o pacote).
// Função main:

// urlConexao: String contendo os detalhes da conexão com o banco de dados MySQL.
// sql.Open("mysql", urlConexao): Abre uma conexão com o banco de dados MySQL usando o driver especificado. Retorna um objeto de banco de dados (db) e um possível erro.
// defer db.Close(): Adia o fechamento da conexão até o final da execução da função main.
// db.Ping(): Tenta pingar o banco de dados para verificar se a conexão está ativa. Se houver algum problema, loga o erro e encerra o programa.
// Se a conexão estiver ativa, imprime uma mensagem indicando que a conexão está aberta.
