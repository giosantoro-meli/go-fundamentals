package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -100000000
	fmt.Println(numero)

	var numero2 uint32 = 12000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 124553
	fmt.Println(numero3)

	//alias
	//UINT8 = byte
	var numero4 byte = 124
	fmt.Println(numero4)

	var float1 float32 = 124.56
	fmt.Println(float1)

	var float2 float64 = 124523523.532556
	fmt.Println(float2)

	var str string = "Texto"
	fmt.Println(str)

	//char é um numero no GO
	char := 'b'
	fmt.Println(char)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("erro interno")
	fmt.Println(erro2)
}
