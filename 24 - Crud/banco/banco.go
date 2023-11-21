package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexao com o mysql
)

// Conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	urlConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConexao)

	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
