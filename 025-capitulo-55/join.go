package main

import (
	"fmt"
	"strings"
)

func main() {
	parrafo := "Esto es un parrafo de prueba para el ejemplo de join"
	arreglo := strings.Split(parrafo, " ")
	// Join une un arreglo de cadenas en una sola cadena
	// el primer parametro es el arreglo de cadenas
	// el segundo parametro es el separador
	for _, valor := range arreglo {
		println(valor)
	}

	cadena := strings.Join(arreglo, " ")
	fmt.Println(cadena)

}
