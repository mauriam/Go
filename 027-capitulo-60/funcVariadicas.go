package main

import "fmt"

//funciones variadicas
//func nombreFuncion(parametros ...tipo) tipoRetorno

func sumar(valores ...int) int {
	//valores es un slice de enteros
	//podemos realizar todas las operaciones que se pueden realizar con un slice
	fmt.Println(valores)
	var total int
	for _, valor := range valores {
		total += valor
	}
	return total
}

func main() {
	total := sumar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)
}
