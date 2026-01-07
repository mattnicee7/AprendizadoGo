package main

import "fmt"

func main() {
	go escrever("ola mundo")
	escrever("programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
	}
}
