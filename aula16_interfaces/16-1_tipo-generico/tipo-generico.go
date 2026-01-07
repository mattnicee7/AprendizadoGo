package main

import "fmt"

func generica(interf ...interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string", 1)
	generica(1)
	generica(true)

	fmt.Println(1, 2, "a", 3.14)

	//
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
	}

	fmt.Println(mapa)
}
