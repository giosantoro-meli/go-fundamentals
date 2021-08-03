package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 5 {
		fmt.Println("Incrementando i: ", i)
		i++
		time.Sleep(time.Second / 2)
	}

	println("----------")

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j: ", j)
		time.Sleep(time.Second / 4)
	}

	println("----------")

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	println("----------")

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	println("----------")

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	println("----------")

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
