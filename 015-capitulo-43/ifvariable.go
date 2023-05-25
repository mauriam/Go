package main

import (
	"fmt"
)

func main() {
	if nombre, edad := "pedro", 14; nombre == "Juan" {
		fmt.Println("Hola, ", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d \n", nombre, edad)
	}

	user := make(map[string]string)

	user["mauriam"] = "mauriamortiz@gmail.com"
	user["daniel"] = "danielurbaez@gmail.com"

	correo, ok := user["mauriam"]
	fmt.Println(correo, ok)

	if correo, ok := user["msm"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("(!) No existe el usuario")
	}
}
