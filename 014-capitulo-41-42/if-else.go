/** App para restaurantes
* Descuentos de 10% hasta 100 de consumo
* Descuentos de 20% hasta 200 de consumo
* Descuentos de 30% mayor 200 de consumo
* IGV 19%
 */
package main

import (
	"fmt"
)

func main() {
	var consumo, descuentos float64
	var datosDescuentos string

	//entrada de datos
	fmt.Print("Ingrese el consumo: ")
	fmt.Scanln(&consumo)

	//proceso
	if consumo < 0 {
		fmt.Println("El consumo no puede ser negativo")
		return
	}
	if consumo <= 100 {
		descuentos = consumo * 0.1
		datosDescuentos = "10%"
	} else if consumo <= 200 {
		descuentos = consumo * 0.2
		datosDescuentos = "20%"
	} else {
		descuentos = consumo * 0.3
		datosDescuentos = "30%"
	}
	// operaciones
	montoDescuento := consumo - descuentos
	igv := montoDescuento * 0.19
	total := montoDescuento + igv

	//salida de datos
	fmt.Println("----------factuta de consumo----------")
	fmt.Println("Descuento que aplica: ", datosDescuentos)
	fmt.Println("Consumo: ", consumo)
	fmt.Println("Descuento: ", descuentos)
	fmt.Println("IGV: ", consumo*0.19)
	fmt.Println("Total: ", total)
}
