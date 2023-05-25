package main

import "fmt"

//funciones variadicas
//func nombreFuncion(parametros ...tipo) tipoRetorno

func sumar(valores ...int) (int, float64) {
	//valores es un slice de enteros
	//podemos realizar todas las operaciones que se pueden realizar con un slice
	fmt.Println(valores)
	var total int
	for _, valor := range valores {
		total += valor
	}
	media := float64(total) / float64(len(valores))
	return total, media
}

func main() {
	total, media := sumar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Suma Total: ", total)
	fmt.Println("Media: ", media)
}
