package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

type entrega struct {
	code     string
	endereco endereco
}

type endereco struct {
	logradouro  string
	numero      string
	complemento string
	cep         string
}

func main() {
	fmt.Println("Arquivo structs")

	var u1 usuario
	u1.nome = "Davi"
	u1.idade = 18

	u2 := usuario{"Giovanni", 28}

	u3 := usuario{nome: "João"}

	u4 := usuario{idade: 23}

	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3.nome)
	fmt.Println(u4.idade)

	en1 := entrega{"1L", endereco{"rua direita", "12", "A", "03300-555"}}

	ad1 := endereco{"rua esquerda", "13", "", "03300-556"}
	en2 := entrega{"2L", ad1}
	fmt.Println(en1)
	fmt.Println(en2)

}
