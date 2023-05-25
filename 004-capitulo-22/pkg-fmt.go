package main

import "fmt"

func main() {
	// fmt.Println("Hello, playground")
	hola := "Hola"
	mundo := "Mundo"
	//sin saltos de lineas
	fmt.Print(hola)
	fmt.Print(mundo)
	//con saltos de lineas
	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Mauriam"
	edad := 32
	//especificando el tipo de dato
	//%s string
	//%d entero
	fmt.Printf("Mi nombre es %s y tengo %d años\n", nombre, edad)
	//otra fomra de hacerlo
	//%v es el valor
	fmt.Printf("Mi nombre es %v y tengo %v años\n", nombre, edad)

	//generando un string y asignandolo a una variable
	mensaje := fmt.Sprintf("Mi nombre es %v y tengo %v años\n", nombre, edad)
	fmt.Println(mensaje)

	//para saber el tipo de dato
	fmt.Printf("El tipo de dato de nombre es %T\n", nombre)
	fmt.Printf("El tipo de dato de edad es %T\n", edad)

	//para ingresar valore por teclado
	var nombre2 string
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scanln(&nombre2)
	fmt.Printf("Hola %s\n", nombre2)
}
