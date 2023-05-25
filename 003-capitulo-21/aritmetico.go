package main

import "fmt"

func main() {
	a := 20
	b := 10
	fmt.Println("Valores iniciales")
	fmt.Println(a)
	fmt.Println(b)

	resultado := a + b
	fmt.Println("Suma: ", resultado)

	resultado = a - b
	fmt.Println("Resta: ", resultado)

	resultado = a * b
	fmt.Println("Multiplicación: ", resultado)

	var div float64 = 3.0 / 2.0
	fmt.Println("División: ", div)

	resultado = a % b
	fmt.Println("Módulo: ", resultado)

	a++
	fmt.Println("Incremento: ", a)

	a--
	fmt.Println("Decremento: ", a)

}
