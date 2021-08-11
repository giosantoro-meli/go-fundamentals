package main

import "fmt"

//funcao recover() recupera a execucao
func recuperaExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {

	defer recuperaExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//Função panic mata a execução do programa
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println("Panic e Recover")
	fmt.Println(alunoEstaAprovado(6, 7))
	fmt.Println("-----------")
	fmt.Println(alunoEstaAprovado(5, 7))
}
