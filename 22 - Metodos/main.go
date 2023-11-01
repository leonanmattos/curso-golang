package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Sprintf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 23}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println(usuario1.idade)
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)

}
