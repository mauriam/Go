package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var resultado int
	for _, valor := range valores {
		num, error := strconv.Atoi(valor)
		if error != nil {
			fmt.Println(error)
			fmt.Println("(!) Error en la expresion")
			fmt.Println("(!) La expresion debe ser de la forma: 1+2+3+4+5+6+7+8+9+10")
			fmt.Println("(!) no puede contener letras")
			return 0
		}
		resultado += num
	}
	return resultado
}

func main() {
	var expresion string
	fmt.Print("=>")
	fmt.Scanln(&expresion)
	fmt.Println("la suma es:", sumar(expresion))
}
