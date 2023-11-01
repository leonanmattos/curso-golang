package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 3 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25

	soma2 := int32(numero1) + numero2

	fmt.Println(soma2)

	// FIM DOS ARITMETICOS

	// ATRIBUICAO

	var variavel1 string = "String"
	variavel2 := "String 2"

	fmt.Println(variavel1, variavel2)

	// FIM ATRIBUICAO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	// FIM OPERADORES RELACIONAIS

	// OPERADORES LOGICOS
	fmt.Println(!false || false)
	// FIM OPERADORES LOGICOS

	// OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 10
	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	// FIM OPERADORES UNARIOS

}
