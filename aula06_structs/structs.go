package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("structs")

	var u usuario
	u.nome = "Matheus"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"rua 1", 123}

	usuario2 := usuario{"uuu", 22, enderecoExemplo}
	fmt.Println(usuario2)

	// criar instaNcia com so um dado (referenciar nome do campo)
	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)

	usuario4 := usuario{nome: "Matheus"}
	fmt.Println(usuario4)

	usuario5 := usuario{"Matheus", 21, endereco{"Rua 3", 123}}

	fmt.Println(usuario5)
}
