package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1 int8, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	resultado := somar(1, 2)
	fmt.Println(resultado)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado2 := f("Texto da funcao")
	fmt.Println(resultado2)

	resultado3, resultado4 := calculosMatematicos(20, 10)
	fmt.Println(resultado3, resultado4)
}
