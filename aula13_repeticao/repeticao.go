package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println(i)
	}

	fmt.Println("Processo finalizado")

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	nomes := []string{"Joao", "Davi", "Lucas"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Arthur",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// while true
	for {
		fmt.Println("infinito")
	}

	// n pode usar range em um struct, somente map array etc
}
