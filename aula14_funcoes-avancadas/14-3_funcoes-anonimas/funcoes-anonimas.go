package main

import "fmt"

func main() {

	func(texto string) {
		fmt.Println(texto)
	}("Ola Mundo")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("oi")

	fmt.Println(retorno)
}
