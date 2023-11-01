package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando")
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando")
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Joao", "Diego", "Maria"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRAS" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":  "Diego",
		"idade": "25",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
