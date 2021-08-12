package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//bom usar ponteiro para alterar um atributo de um struct
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	fmt.Println("Usuário é maior de idade? ", usuario1.maiorDeIdade())

	usuario1.fazerAniversario()

	fmt.Printf("Usuário %s fez aniversário! Sua nova idade é %d anos\n", usuario1.nome, usuario1.idade)
}
