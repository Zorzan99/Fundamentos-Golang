// O pacote "servidor" contém funcionalidades relacionadas à lógica do servidor HTTP.

package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Definição da estrutura do usuário para codificação e decodificação JSON.
type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um novo usuário no banco de dados.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// Leitura do corpo da requisição.
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	// Declaração de uma variável do tipo "usuario" para armazenar os dados do JSON.
	var usuario usuario

	// Decodificação do JSON para a estrutura do usuário.
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro na decodificação do JSON: "))
		return
	}

	// Conexão com o banco de dados.
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	// Preparação da declaração SQL para inserir um novo usuário.
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	// Execução da declaração SQL para inserção do usuário.
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	// Obtenção do ID do usuário inserido.
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter ID inserido"))
		return
	}

	// Resposta indicando sucesso e o ID do usuário inserido.
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

// BuscarUsuarios retorna todos os usuários salvos no banco de dados.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Conexão com o banco de dados.
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	// Execução de uma consulta SQL para obter todos os usuários.
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	// Declaração de uma slice para armazenar os usuários.
	var usuarios []usuario

	// Iteração sobre as linhas retornadas pela consulta SQL.
	for linhas.Next() {
		var usuario usuario

		// Leitura dos valores das colunas da tabela e atribuição à estrutura do usuário.
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao buscar usuário"))
			return
		}

		// Adição do usuário à slice de usuários.
		usuarios = append(usuarios, usuario)
	}

	// Resposta com o status HTTP OK e a lista de usuários convertida para JSON.
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuários para JSON"))
		return
	}
}

// BuscarUsuario retorna um usuário específico salvo no banco de dados.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// Obtenção dos parâmetros da rota usando o mux.
	parametros := mux.Vars(r)

	// Conversão do ID para um número inteiro.
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro na conversão do ID"))
		return
	}

	// Conexão com o banco de dados.
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro na conexão com o banco de dados"))
		return
	}

	// Execução de uma consulta SQL para obter um usuário específico pelo ID.
	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	// Declaração de uma variável do tipo "usuario" para armazenar o usuário encontrado.
	var usuario usuario

	// Leitura do usuário encontrado e atribuição à estrutura do usuário.
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	// Resposta com o status HTTP OK e o usuário convertido para JSON.
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para JSON"))
		return
	}
}

// AtualizarUsuario atualiza um usuário no banco de dados com base no ID fornecido.
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// Obtenção dos parâmetros da rota usando o mux.
	parametros := mux.Vars(r)

	// Conversão do ID para um número inteiro.
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro na conversão do ID"))
		return
	}

	// Leitura do corpo da requisição.
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro na leitura do corpo da requisição"))
		return
	}

	// Declaração de uma variável do tipo "usuario" para armazenar os dados do JSON.
	var usuario usuario

	// Decodificação do JSON para a estrutura do usuário.
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro na conversão do JSON para usuário"))
		return
	}

	// Conexão com o banco de dados.
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro na conexão com o banco"))
		return
	}
	defer db.Close()

	// Preparação da declaração SQL para atualizar um usuário pelo ID.
	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	// Execução da declaração SQL para atualização do usuário.
	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	// Resposta indicando sucesso sem conteúdo.
	w.WriteHeader(http.StatusNoContent)
}

// DeletarUsuario remove um usuario do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro na conversão do ID"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro na conexão com o banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
