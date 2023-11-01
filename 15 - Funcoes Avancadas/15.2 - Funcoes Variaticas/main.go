package main

import "fmt"

func soma(numeros ...int) int {
	s := 0
	for _, v := range numeros {
		s += v
	}
	return s
}

func escrever(texto string, numeros ...int) {
	for numero := range numeros {
		fmt.Println(texto, numero)
	}
}
func main() {
	total := soma(1, 2, 3, 4, 5, 4, 5, 3, 3, 4, 5)
	fmt.Println(total)

	escrever("Leonan", 1, 2, 3, 4, 5, 6)
}
