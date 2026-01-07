package main

import "fmt"

var n int

// executa antes da main
// uma por arquivo (nao por pacote)
func init() {
	n = 10
	fmt.Println("executando o init")
}

func main() {
	fmt.Println("main")
	fmt.Println(n)
}
