package main

import "fmt"

// go nao deixa usar variavel sem ser utilizada, import
func main() {
	// tipos explicitos
	var nome string = "matt"
	fmt.Println(nome)

	// no tip (inferencia de tipo baseado no valor)
	nome2 := "matt 2" // (string) implicita
	fmt.Println(nome2)

	var (
		variavel3 string = "teste"
		variavel4 string = "teste"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6, variavel7 := "5", "6", "7"

	fmt.Println(variavel5, variavel6, variavel7)

	// constantes
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// mudanca de valor
	variavel5 = "teste 5"
	fmt.Println(variavel5)

	// mudanca
	variavel6, variavel7 = variavel7, variavel6

	fmt.Println(variavel6, variavel7)
}
