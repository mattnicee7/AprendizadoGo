package main

import "fmt"

func main() {

	// aritmeticos
	// + - / * %
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// nao pode utilizar variaveis com tipo diferentes (int16 + int32, etc)
	//	var numero1 int16 = 10
	//	var numero2 int32 = 25
	//soma2 := numero1 + numero2 (ERRO)

	var numero1, numero2 int16 = 10, 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	var numero11 int16 = 10
	var numero22 int16 = 25
	soma3 := numero11 + numero22
	fmt.Println(soma3)

	// tipos atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// fim operadores atribuicao

	// relacionais (boolean)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// fim

	fmt.Println("---------------")

	// operadores logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// fim operadores logicos

	fmt.Println("---------------")

	// operadores unarios
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	//--numero nao existe (incrementar depois de verificar)
	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 3
	fmt.Println(numero)
	numero %= 2
	fmt.Println(numero)
	// fim operadores unarios

	// operador ternario
	// se fosse em outra lang:
	//texto := numero > 5 ? "maior que 5" : "menor que 5"
	// NAO EXISTE EM GOLANG, somente com if else

	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)

}
