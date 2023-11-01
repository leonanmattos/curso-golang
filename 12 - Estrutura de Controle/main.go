package main

import "fmt"

func main() {
	numero := 10

	if numero > 5 {
		fmt.Println("O número é maior que 5")
	} else {
		fmt.Println("O número é menor ou igual a 5")
	}

	if outroNumero := numero; outroNumero > 9 {
		fmt.Println("O número é maior que 9")
	} else {
		fmt.Println("Não vai entrar nesse bloco")
	}
}
