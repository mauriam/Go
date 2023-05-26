package main

import "fmt"

func main() {
	// Funcion anonima
	func() {
		fmt.Println("Hola mundo desde una funcion anonima")
	}()
	// Funcion anonima asignada a una variable
	myfuncion := func() {
		fmt.Println("Hola mundo desde una funcion anonima asignada a una variable")
	}
	myfuncion() // Llamada a la funcion anonima asignada a una variable

	// Funcion anonimas asignada a un arreglo
	funciones := []func(){
		func() {
			fmt.Println("funcion indice cero")
		},
		func() {
			fmt.Println("funcion indice uno")
		},
		func() {
			fmt.Println("funcion indice dos")
		},
	}

	funciones[0]() // Llamada a la funcion anonima en el indice cero
	funciones[2]() // Llamada a la funcion anonima en el indice dos

	// Funcion anonima con parametros
	func(nombre string) {
		fmt.Println("Hola", nombre)
	}("Golang")

	// Funcion anonima con retorno
	funciones2 := []func(nombre string){
		func(nombre string) {
			fmt.Println("Hola", nombre)
		},
		func(nombre string) {
			fmt.Println("Hola un gusto en saludarte", nombre)
		},
		func(nombre string) {
			fmt.Println("Hola como estas", nombre)
		},
	}
	funciones2[2]("mauriam")
	funciones2[1]("mauriam")
}
