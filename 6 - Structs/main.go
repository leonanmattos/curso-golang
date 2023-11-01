package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivos Structs")

	var u usuario = usuario{
		nome:  "Joao",
		idade: 25,
	}

	enderecoExemplo := endereco{"Rua dos Lobos", 123}

	fmt.Println(u.nome)
	fmt.Println(u.idade)

	u2 := usuario{
		nome:     "Maria",
		idade:    30,
		endereco: enderecoExemplo,
	}

	fmt.Println(u2)
}
