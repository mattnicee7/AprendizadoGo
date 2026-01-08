package main

import (
	"AprendizadoGo/aula19_testes-automatizados/19-1_introducao/enderecos"
	"fmt"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
