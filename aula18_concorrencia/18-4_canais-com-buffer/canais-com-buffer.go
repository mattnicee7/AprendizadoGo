package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Ola Mundo!"
	canal <- "PRogramando Go"
	canal <- "terceiro valor"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
