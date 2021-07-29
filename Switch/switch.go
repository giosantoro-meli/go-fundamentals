package main

import "fmt"

func main() {
	fmt.Println("Switch")
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(4))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(6))
	fmt.Println(diaDaSemana(7))
	fmt.Println(diaDaSemana(8))
}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Valor inválido"
	}
}
