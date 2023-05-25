/*

Practica 02: Calcular cociente y el Residuo de dos números enteros

Enunciado: halar el cociente y el residuo (resto) de dos números enteros.

Análisis: para la solución de este problema, se requiere que el usuario ingrese dos números enteros y el sistema realice el cálculo respectivo para hallar el cociente y residuo.

*/

package main

import "fmt"

func main() {
	var numero_uno, numero_dos int
	fmt.Println("Ingrese el primer número: ")
	fmt.Scanln(&numero_uno)
	fmt.Println("Ingrese el segundo número: ")
	fmt.Scanln(&numero_dos)
	fmt.Println("El cociente de los dos números es: ", numero_uno/numero_dos)
	fmt.Println("El residuo de los dos números es: ", numero_uno%numero_dos)

}
