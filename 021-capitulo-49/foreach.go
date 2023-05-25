package main

import "fmt"

func main() {
	nombres := [...]string{"Alejandro", "Maria", "Pedro", "Pablo", "Diana", "Eduardo"}
	// itereacion de un array con for tradicional
	for i := 0; i < len(nombres); i++ {
		println(nombres[i])
	}
	// itereacion de un array con foreach
	for indice, valores := range nombres {
		fmt.Println("Indice:", indice, "Valor:", valores)
	}
}
