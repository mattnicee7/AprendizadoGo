package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("maior q 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("maior q 0")
	} else if numero < -10 {
		fmt.Println("menor q -10")
	} else {
		fmt.Println("esta entre 0 e -10")
	}
}
