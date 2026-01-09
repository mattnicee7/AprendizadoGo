package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // se colocar um - ele nao vai aparecer no json
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Rex","raca":"Dalmata","idade":3}`

	var c cachorro
	// endereco de memoria que vai jogar
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJson := `{"nome": "Toby", "raca": "Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
