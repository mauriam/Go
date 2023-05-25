package main

import "fmt"

func main() {
	var numeros [10]int
	//modificar el valor de un elemento del array
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)    //imprimr todo el array
	fmt.Println(numeros[2]) //imprimir un elemento del array

	//declarar un array con valores
	var nombres = [3]string{"Juan", "Pedro", "Luis"}
	fmt.Println(nombres)

	//declarar un array sin saber longitud
	var colores = [...]string{"Rojo", "Verde", "Azul", "azul"}
	fmt.Println(colores)

	//imprimir la longitud del array
	fmt.Println(len(colores))

	monedas := [...]string{
		0: "peso",
		1: "dolar",
		7: "euro",
	}
	fmt.Println(monedas)
	fmt.Println(monedas[7])
	fmt.Println(len(monedas))
	// sub array
	fmt.Println("sub array")
	fmt.Println(monedas[0:2])
}
