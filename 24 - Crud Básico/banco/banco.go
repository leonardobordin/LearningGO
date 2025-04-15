package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conexao() (*sql.DB, error) {
	stringConexao := "leonardo:Leonardo@2805@/devbook?charset=utf8&parseTime=true&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
