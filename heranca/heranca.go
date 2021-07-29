package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança que não é bem herança")

	p1 := pessoa{"João", "Silva", 25, 178}
	e1 := estudante{p1, "Biologia", "USP"}

	fmt.Println(e1)
	fmt.Println("Nome completo do estudante:", e1.nome, e1.sobrenome)
	fmt.Println("Altura do estudante:", e1.altura, "cm")
}
