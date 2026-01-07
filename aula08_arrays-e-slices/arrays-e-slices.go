package main

import (
	"fmt"
)

func main() {
	var array1 [5]string
	array1[0] = "pos1"
	array1[1] = "pos2"
	array1[2] = "pos3"
	array1[3] = "pos4"
	array1[4] = "pos5"
	fmt.Println(array1)

	array2 := [5]string{"pos1", "pos2", "pos3", "pos4", "pos5"}
	fmt.Println(array2)

	// vao fixar o tamanho do array de acordo com a quantidade de itens que foi passado
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18, 19, 20, 21, 22)

	fmt.Println(slice)

	// criar um slice de um array, da pos1 ate pos3.
	slice2 := array2[1:3] // indice1 = inclusivo // indice3 = exclusivo
	// nesse caso vai so pegar a posicao 1 e 2, pq o 3 e exclusivo
	fmt.Println(slice2)

	// ele n cria uma copia, ele so aponta para as posicoes originais do array
	// se eu alterar no array, o slice tb vai ser alterado
	array2[1] = "Pos alterada"
	fmt.Println(slice2)
}
