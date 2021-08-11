package main

import "fmt"

func funcao1() {
	fmt.Println("Executando func 1")
}

func funcao2() {
	fmt.Println("Executando func 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	return (n1+n2)/2 >= 6
}

func main() {
	fmt.Println("Método defer")
	//DEFER = ADIAR
	defer funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(5, 7))
}

//DEFER é útil para abrir/fechar conexões com o banco de dados
