package main

import "fmt"

var global string

func primera_funcion() {
	global = "Primera Funcion"
	println(global)
}
func segunda_funcion() {
	global = "Segunda Funcion"
}
func tercera_funcion() {
	global = "Tercera Funcion"
}

func main() {
	global = "Hola Mundo"
	primera_funcion()
	segunda_funcion()
	fmt.Println(global)
	tercera_funcion()
	fmt.Println(global)
}
