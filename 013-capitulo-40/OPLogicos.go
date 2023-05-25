package main

import "fmt"

func main() {
	// Not
	// !, si es verdadero, el resultado es falso

	// And
	// &&, si ambos son verdaderos, el resultado es verdadero

	// Or
	// ||, si alguno es verdadero, el resultado es verdadero

	//NOT ejemplo
	fmt.Println("\nEjemplo NOT !: si es verdadero, el resultado es falso")
	fmt.Println("!true  ", !true)

	//AND ejemplo
	fmt.Println("\n Ejemplo AND &&: si ambos son verdaderos, el resultado es verdadero")
	fmt.Println("true  && true  ", true && true)
	fmt.Println("true  && false ", true && false)
	fmt.Println("false && true  ", false && true)
	fmt.Println("false && false ", false && false)

	//OR ejemplo
	fmt.Println("\n Ejemplo OR ||: si alguno es verdadero, el resultado es verdadero")
	fmt.Println("true  || true  ", true || true)
	fmt.Println("true  || false ", true || false)
	fmt.Println("false || true  ", false || true)
	fmt.Println("false || false ", false || false)

	//ejemplo complejo
	b := 2
	r := b == 2 || (4 > b)
	fmt.Println(r)
}
