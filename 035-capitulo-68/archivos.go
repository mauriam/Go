package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a file
	//file, err := os.Create("test.txt")

	// Open a file
	file, err := os.Open("hola.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("[!] error al abrir el archivo [!]")
	} else {
		contenido := make([]byte, 254)
		// file.Read(contenido) // lee el archivo y lo guarda en contenido
		// devuelve la cantidad de bytes leidos y un error
		longitud, _ := file.Read(contenido)
		contenido_archivo := string(contenido[:longitud])
		fmt.Println("Longitud: ", longitud)
		fmt.Println("Contenido: \n", contenido_archivo)
	}

	// Write a file
	// Read a file
	// Close a file
}
