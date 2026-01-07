package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("salvando dados do usuario %s no banco de dados", u.nome)
	fmt.Println()
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// alterar variaveis
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"davi", 30}
	usuario2.salvar()

	fmt.Println(usuario2.maiorDeIdade())
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
