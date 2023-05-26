package main

import "fmt"

// Los punteros son variables que almacenan la direccion de memoria de otra variable
// Los punteros se declaran con el simbolo *
// Los punteros se acceden con el simbolo *
// Los punteros se acceden con el simbolo &

func modificar(c *string) {

	// accediendo a la direccion de memoria
	*c = "cadena modificada"
}

func main() {
	// declarando una variable
	cadena := "hola mundo"

	// mostrando cadena sin modificar
	fmt.Println("antes de la funcion: ", cadena)

	// llamando a la funcion modificar y pasando la direccion de memoria de la variable cadena
	modificar(&cadena)

	// mostrando cadena aqui ya fue modificada por la funcion
	fmt.Println("antes de la funcion: ", cadena)

}
