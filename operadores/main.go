package main

import "fmt"

func main() {
	//ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 1.0 / 2
	restoDaDivisao := 1 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	//ATRIBUICAO
	var string1 string = "String 1"
	string2 := "String 2"
	fmt.Println(string1, string2)

	//RELACIONAIS
	bool1 := (1 > 2)
	bool2 := (1 >= 2)
	bool3 := (1 < 2)
	bool4 := (1 <= 2)
	bool5 := (1 == 2)
	bool6 := (1 != 2)
	fmt.Println(bool1, bool2, bool3, bool4, bool5, bool6)

	//LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso, verdadeiro || falso, !verdadeiro, !falso)

	//UNARIOS
	numero := 10
	numero++
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero *= 2
	fmt.Println(numero)

	//OPERACAO TERNARIA NAO EXISTE NO GO

}
