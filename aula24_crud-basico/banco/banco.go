package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver Conexao
)

// Conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "root:123456@tcp(localhost:3306)/mattnicee7?charset=utf8&parseTime=true&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
