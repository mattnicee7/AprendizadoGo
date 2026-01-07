package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Ola Mundo!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("PRogramando em Go")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Goroutine 3")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 4")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
