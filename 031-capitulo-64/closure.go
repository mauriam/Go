package main

import "strings"

// Funcion que retorna una funcion
// los parametros de la funcion retornada son los parametros de la funcion que retorna
// los parametros de la funcion que retorna son los parametros de la funcion que retorna

func repetir(n int) func(string) string {
	return func(texto string) string {
		return strings.Repeat(texto, n)
	}
}

func main() {
	// Funcion que retorna una funcion
	repetirCincoVeces := repetir(5)
	// Funcion que retorna una funcion
	repetirDiezVeces := repetir(10)

	println(repetirCincoVeces("Hola "))
	// Hola Hola Hola Hola Hola
	println(repetirDiezVeces("Mundo "))
	// Mundo Mundo Mundo Mundo Mundo Mundo Mundo Mundo Mundo Mundo
}
