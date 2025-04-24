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

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Falha ao converter json"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao se conectar com a base de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES(?, ?)")
	if erro != nil {
		w.Write([]byte("Falha ao criar statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Falha ao executar statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Falha ao buscar Id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso. Id: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao se conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		w.Write([]byte("Falha ao buscar os usuarios"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Falha ao escanear usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Falha ao converter para JSON"))
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha ao converter ID"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao se conectar com a base de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Falha ao buscar ID selecionado"))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Falha ao converter para JSON"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha ao converter ID"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Falha ao converter o corpo da requisição para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha na conexao com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Falha ao preparar statement"))
	}
	defer statement.Close()

	if _, erro := statement.Exec(&usuario.Nome, &usuario.Email, ID); erro != nil {
		w.Write([]byte("Falha ao atualizar o usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha ao converter ID"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao se conectar a base de dados"))
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Falha ao preparar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Falha ao excluir usuario"))
	}

	w.WriteHeader(http.StatusNoContent)
}
