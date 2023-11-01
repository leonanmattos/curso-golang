package main

import (
	"errors"
	"fmt"
)

func main() {

	// NUMEROS INTEIROS
	var numero int8 = 100
	fmt.Println(numero)

	var numero2 uint = 100
	fmt.Println(numero2)

	// alias
	// INT32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	// FIM NUMEROS INTEIROS
	// NUMEROS REAIS

	var numero4 byte = 123
	fmt.Println(numero4)

	var numero5 float32 = 321321321321231321231231231231321231232.3232232332323232
	fmt.Println(numero5)

	var numero6 float64 = 332121323121323123112321323123123123121323121323121323123121313332323232323232323232323123123123123231.32
	fmt.Println(numero6)

	// FIM NUMEROS REAIS

	// STRING

	var str1 string = "Texto"
	fmt.Println(str1)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'P'
	fmt.Println(char)

	// FIM STRINGS

	var texto int16
	fmt.Println(texto)

	var booleano1 bool = false
	fmt.Println(booleano1)

	var erro error = errors.New("Erro ao processar")
	fmt.Println(erro)

}
