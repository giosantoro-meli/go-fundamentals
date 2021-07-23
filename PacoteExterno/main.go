package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello")

	erro := checkmail.ValidateFormat("devbookgmail.com")
	fmt.Println(erro)
}
