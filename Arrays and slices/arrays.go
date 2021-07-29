package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and slices")

	var arr1 [5]string
	arr1[0] = "Cat"
	arr1[1] = "Dog"
	arr1[2] = "Bird"
	arr1[3] = "Fish"
	arr1[4] = "Turtle"
	fmt.Println(arr1)

	arr2 := [5]string{"Posicao1", "Posicao2", "Posicao3", "Posicao4", "Posicao5"}
	fmt.Println(arr2)

	//Slices não tem tamanho limitado
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	fmt.Println(slice)

	slice = append(slice, 19)
	fmt.Println(slice)

	//Pegando fatias de um array (funciona como um ponteiro)
	slice2 := arr1[0:2]
	fmt.Println(slice2)

	//Apesar de parecidos, são diferentes
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr1))

	//=====================//

	//Arrays Internos
	fmt.Println("-----------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println("Length: ", len(slice3))
	fmt.Println("Capacity: ", cap(slice3))
}
