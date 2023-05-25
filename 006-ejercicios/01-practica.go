/*
Practica 01: Suma de dos números enteros
Enunciado: dado dos números enteros, hallar la suma.
Análisis: para la solución de este problema, se requiere que el usuario ingrese dos números enteros y el sistema realice el cálculo respectivo para hallar la suma.
*/
package main

import "fmt"

func main() {
	var numero_uno, numero_dos int
	fmt.Println("Ingrese el primer número: ")
	fmt.Scanln(&numero_uno)
	fmt.Println("Ingrese el segundo número: ")
	fmt.Scanln(&numero_dos)
	fmt.Println("La suma de los dos números es: ", numero_uno+numero_dos)

}
