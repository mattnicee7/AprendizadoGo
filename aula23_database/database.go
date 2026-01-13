package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "root:123456@tcp(localhost:3306)/mattnicee7?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", stringConexao)

	if err != nil {
		fmt.Println("Dentro sql")
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("Dentro ping")
		log.Fatal(err)
	}

	fmt.Println("Conexao aberta")

	linhas, err := db.Query("select * from mattnicee7")
	if err != nil {
		log.Fatal(err)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
