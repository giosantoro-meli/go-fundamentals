package main

import "fmt"

var n int

func init() {
	fmt.Println("Função init sendo executada")
	n = 10
}

func main() {
	fmt.Println("Função principal sendo executada")
	fmt.Println(n)
}
