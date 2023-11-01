package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func funcao3(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado sera retornado")
	fmt.Println("Entrando na funcao para verificar se aluno esta aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// funcao1()
	// funcao2()

	fmt.Println(funcao3(7, 8))
}
