package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go func() {
		escrever("Ola mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 6; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
