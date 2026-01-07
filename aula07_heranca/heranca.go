package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Joao", "Pedro", 20, 170}
	fmt.Println(p1)

	// tem que passar uma pessoa como campo
	e1 := estudante{p1, "medicina", "usp"}
	fmt.Println(e1.altura)
	fmt.Println(e1.nome)
	fmt.Println(e1.faculdade)

	e2 := estudante{pessoa{"matheus", "matheus", 21, 170}, "medicina", "usp"}
	fmt.Println(e2)
}
