package main

import "strings"

func main() {
	cadena := "esto es una cadena de prueba"

	// Split separa una cadena en un arreglo de cadenas
	// el primer parametro es la cadena a separar
	// el segundo parametro es el separador
	// el tercer parametro es la cantidad de elementos del arreglo
	// que se van a crear
	// si se pone -1 se crean todos los elementos posibles
	// si se pone 0 no se crea ningun elemento
	// si se pone 1 se crea un elemento
	// si se pone 2 se crean dos elementos
	// etc
	// si no se pone el tercer parametro se crean todos los elementos posibles
	arreglo := strings.Split(cadena, " ")
	for _, valor := range arreglo {
		println(valor)
	}

}
