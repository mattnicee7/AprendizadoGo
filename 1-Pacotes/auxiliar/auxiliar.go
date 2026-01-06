package auxiliar

import "fmt"

// se uma funcao eh toda minuscula, ela so vai ser visivel dentro do pacote que ela esta
// maiuscula = pode ser exportada por outros pacotes.

func Escrever() {
	fmt.Println("Arquivo Auxiliar")
	escrever2()
}
