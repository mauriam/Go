package main

import "fmt"

func main() {
	a := 10
	b := 10
	var r bool

	//operadores relacionales >, <, >=, <=, ==, !=
	// igualdad
	r = a == b
	// usamos %t para imprimir un booleano
	fmt.Printf("\n %d es igual que %d? R: %t", a, b, r)
	// desigualdad
	r = a != b
	fmt.Printf("\n %d es diferente que %d? R: %t", a, b, r)
	// mayor que
	r = a > b
	fmt.Printf("\n %d es mayor que %d? R: %t", a, b, r)
	// menor que
	r = a < b
	fmt.Printf("\n %d es menor que %d? R: %t", a, b, r)
	// mayor o igual que
	r = a >= b
	fmt.Printf("\n %d es mayor o igual que %d? R: %t", a, b, r)
	// menor o igual que
	r = a <= b
	fmt.Printf("\n %d es menor o igual que %d? R: %t", a, b, r)

}
