package main

import "fmt"

func main() {
	fmt.Println("Funções anônimas")

	func() {
		fmt.Println("Olá Mundo")
	}()

	func(n1 int, n2 int) {
		fmt.Println(n1 + n2)
	}(1, 2)
}
