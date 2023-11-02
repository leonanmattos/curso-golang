package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola mundo", canal)

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 6; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
