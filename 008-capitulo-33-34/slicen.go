package main

import "fmt"

func main() {
	//Slicen
	numeros := []int{10, 20, 30, 40, 50}

	fmt.Println(numeros, len(numeros), cap(numeros))

	//agregar elemento al slicen
	numeros = append(numeros, 60)
	fmt.Println(numeros, len(numeros), cap(numeros))

	//sub slicen
	subslice := numeros[0:2]
	fmt.Println(subslice)

}
