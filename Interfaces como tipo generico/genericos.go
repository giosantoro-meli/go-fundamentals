package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Tipo genérico")
	generica("String")
	generica(1)
	generica(true)
}
