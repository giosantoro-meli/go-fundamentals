package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
		"estudo": {
			"curso":   "Engenharia",
			"periodo": "noturno",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["estudo"])
	fmt.Println(usuario2["nome"]["primeiro"])

	usuario3 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	delete(usuario3, 3) //map, chave a ser deletada

	fmt.Println(usuario3)

	usuario3[4] = "D"

	fmt.Println(usuario3) //adicionando chaves ao map
}
