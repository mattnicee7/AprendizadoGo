package main

import (
	"AprendizadoGo/1-Pacotes/auxiliar"
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("matt@gmail.com")
	fmt.Println(erro)
}
