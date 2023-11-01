package main

import "fmt"

func calculosMatematicos(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	v, c := calculosMatematicos(10, 15)
	fmt.Println(v, c)

}
