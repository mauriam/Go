package main

import "fmt"

func main() {
	var vocal string
	fmt.Print("Ingresa una letra: ")
	fmt.Scanf("%s", &vocal)

	switch vocal {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("Es una vocal")
	default:
		fmt.Println("No es una vocal")
	}
}
