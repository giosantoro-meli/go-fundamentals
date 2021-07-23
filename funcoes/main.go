package main

import "fmt"

func somar(n1 int32, n2 int32) int32 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int32) (int32, int32) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	fmt.Println(calculosMatematicos(10, 5))

	//o _ permite ignorar retornos da função
	resultadoSoma, _ := calculosMatematicos(10, 5)
	fmt.Println(resultadoSoma)
}
