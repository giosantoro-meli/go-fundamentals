package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

//apenas um variático por função, sempre na última posição
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções variáticas")
	x0 := soma(1, 2, 3, 4, 5, 6, 200, 12, 13)
	println(x0)

	escrever("Olá Mundo", 1, 2, 3, 4, 5, 6, 7)
}
