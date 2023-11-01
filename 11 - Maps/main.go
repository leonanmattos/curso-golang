package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":  "Diego",
		"idade": "25",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"sobrenome": "sobrenome",
			"idade":     "25",
		},
		"nome2": {
			"sobrenome": "sobrenome",
			"idade":     "25",
		},
		"nome3": {
			"sobrenome": "sobrenome",
			"idade":     "25",
		},
	}

	delete(usuario2, "nome2")

	fmt.Println(usuario2)
}
