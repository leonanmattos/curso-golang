package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "SEGUNDA-FEIRA"
	case 2:
		return "TERÇA-FEIRA"
	default:
		return "NÚMERO INVÁLIDO"
	}
}

func main() {
	diaSemanaStr := diaSemana(4)
	fmt.Println(diaSemanaStr)
}
