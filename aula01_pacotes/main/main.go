package main

import (
	"AprendizadoGo/aula01_pacotes/auxiliar"
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("matt@gmail.com")
	fmt.Println(erro)
}
