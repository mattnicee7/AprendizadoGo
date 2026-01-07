package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	// quando atribui valor a uma variavel, o valor eh uma copia
	// se eu atribuo o valor de uma variavel e coloco em outra, ele cria uma copia
	// entao se eu alterar a variavel1, o valor da variavel2 nao vai ser alterado

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro = referencia de memoria
	// Ponteiro sem valor = <nil>
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // desreferenciacao
}
