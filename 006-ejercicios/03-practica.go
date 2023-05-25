/*
Practica 03: Calcular el Precio de Venta

Enunciado: dado el valor de venta de productos, hallar el IGV (18%) y el precio de venta.

Análisis: para la solución de este problema, se requiere que el usuario ingrese el valor de venta del producto y el sistema realice el cálculo respectivo para hallar el IGV y el precio de venta.

*/

package main

import "fmt"

func main() {
	var precio_venta float32
	fmt.Println("Ingrese el precio de venta: ")
	fmt.Scanln(&precio_venta)
	fmt.Println("El IGV es: ", precio_venta*0.18)
	fmt.Println("El precio de venta es: ", precio_venta+(precio_venta*0.18))
}
