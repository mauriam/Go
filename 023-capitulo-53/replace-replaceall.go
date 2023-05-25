package main

import "strings"

func main() {
	frase := "esto es una frase de prueba que sera encrytada"
	//frase donde replace, lo que se va a reemplazar, por lo que se va a reemplazar, cuantas veces se va a reemplazar
	// -1 es todas las veces que se encuentre
	// 0 es ninguna vez
	// 1 es una vez
	// 2 es dos veces
	frase = strings.Replace(frase, "a", "4", -1)
	//ReplaceAll es lo mismo que Replace con -1 como parametro de veces
	frase = strings.ReplaceAll(frase, "e", "3")
	frase = strings.ReplaceAll(frase, "i", "1")
	frase = strings.ReplaceAll(frase, "o", "0")
	frase = strings.ReplaceAll(frase, "u", "2")

	println(frase)
}
