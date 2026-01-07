package main

import "fmt"

func main() {
	// dentro do colchete = key
	// fora do colchete = value
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}

	fmt.Println(usuario2)

	map2 := map[int8]string{
		1: "matheus",
		2: "arthur",
		3: "teste",
	}

	fmt.Println(map2)
}
