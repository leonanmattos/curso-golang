package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Joao", "Pedro", 20}
	fmt.Println(p1)
	p2 := estudante{p1, "Engenharia de Software", "UFSC"}
	fmt.Println(p2.nome)

}
