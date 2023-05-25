package main

import "fmt"

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	// Adicionando valores
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)

	dias[7] = "SABADO"

	fmt.Println(dias)
	//eliminar un elemento
	delete(dias, 1)

	fmt.Println(dias, len(dias))

}
