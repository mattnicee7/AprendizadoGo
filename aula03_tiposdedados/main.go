package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos inteiros
	//int8, int16, int32, int64
	var numero int64 = 100678134291
	fmt.Println(numero)

	// cria de acordo com os bits do pc (32 ou 64)
	var numero2 int = 132949134
	fmt.Println(numero2)

	numero3 := 18348913
	fmt.Println(numero3)
	//

	// unsigned int
	// n aceita numero negativos
	// uint8, 16, 32, 64
	var numero4 uint32 = 10000
	fmt.Println(numero4)

	// alias (rune = int32)
	var numero5 rune = 123456
	fmt.Println(numero5)

	// alias (byte = int8)
	var numero6 byte = 127
	fmt.Println(numero6)

	// numeros reais (float32 ou float64)
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 137431849.1348134913
	fmt.Println(numeroReal2)

	numeroReal3 := 13741834.18348134
	fmt.Println(numeroReal3)
	//

	var str string = "matt"
	fmt.Println(str)

	name := "matheus"
	fmt.Println(name)

	// go nao tem char
	// traduz da tabela ascii ('')
	char := 'A'
	fmt.Println(char)
	//

	// sem valor inicial
	// string = ""
	// numeros = 0
	// ... = nil
	var text string
	fmt.Println(text)
	//

	var booleano1 bool = true
	fmt.Println(booleano1)

	//

	var erro error = errors.New("Internal Error")
	fmt.Println(erro)
	// <nil> padrao
}
