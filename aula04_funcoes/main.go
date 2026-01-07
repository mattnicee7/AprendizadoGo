package main

import "fmt"

// c retorno
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// retorno duplo
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// declarar funcao como variavel
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	txt := f("oi")
	fmt.Println(txt)

	// chama a funcao, _ ignora o segundo resultado
	resultadoSoma, _ := calculosMatematicos(1, 2)
	fmt.Println(resultadoSoma)

	//
}
