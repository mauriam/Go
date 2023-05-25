package main

import "fmt"

//funcion  sin parametros
func saludar() {
	fmt.Println("Hola Mundo desde una función Saludar")
}

//funcion con parametros
func saludar2(nombre string) {
	fmt.Println("Hola", nombre, "desde una función Saludar2")
}

//funcion con parametros y retorno
func saludar3(nombre string) string {
	mensaje := fmt.Sprintf("Hola %s desde una función Saludar3", nombre)
	return mensaje
}

//funcion con parametros y retorno
func suma(a, b int) int {
	return a + b
}

//funcion Main
func main() {
	saludar()
	saludar2("Mauriam")
	fmt.Println(saludar3("Mauriam"))

	fmt.Println("la suma de 10 y 20 es", suma(10, 20))
}
