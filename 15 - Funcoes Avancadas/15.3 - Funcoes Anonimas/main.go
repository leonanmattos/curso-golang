package main

import "fmt"

func main() {
	func() {
		fmt.Println("Ola mundo")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Texto recebido %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}
