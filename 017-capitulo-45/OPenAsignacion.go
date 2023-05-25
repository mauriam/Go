package main

import "fmt"

func main() {

	a := 2
	a = a + 5
	fmt.Println(a)
	// suma en asignacion
	a += 5
	fmt.Println(a)
	// resta en asignacion
	a -= 5
	fmt.Println(a)
	// multiplicacion en asignacion
	a *= 5
	fmt.Println(a)
	// division en asignacion
	a /= 5
	fmt.Println(a)
	// modulo en asignacion
	a %= 5
	fmt.Println(a)
	// AND en asignacion
	a &= 5
	fmt.Println(a)
	// OR en asignacion
	a |= 5
	fmt.Println(a)
	// XOR en asignacion
	a ^= 5
	fmt.Println(a)
	// Corrimiento a la izquierda en asignacion
	a <<= 5
	fmt.Println(a)
	// Corrimiento a la derecha en asignacion
	a >>= 5
	fmt.Println(a)
}
