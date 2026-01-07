package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terca"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Invalido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemanaString string

	switch {
	case numero == 1:
		diaDaSemanaString = "Domingo"
	case numero == 2:
		diaDaSemanaString = "Segunda"
	case numero == 3:
		diaDaSemanaString = "Terca"
	case numero == 4:
		diaDaSemanaString = "Quarta"
	case numero == 5:
		diaDaSemanaString = "Quinta"
	case numero == 6:
		diaDaSemanaString = "Sexta"
	case numero == 7:
		diaDaSemanaString = "Sabado"
	default:
		diaDaSemanaString = "invalido"
	}

	return diaDaSemanaString
}

func main() {
	fmt.Println("oi")
	dia := diaDaSemana(200)
	fmt.Println(dia)

	dia2 := diaDaSemana2(6)
	fmt.Println(dia2)
}
