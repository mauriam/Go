package main

import (
	"errors"
	"fmt"
)

func div(x, y float64) float64 {
	return x / y
}
func divint(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("No se puede dividir por cero")
	} else {
		return x / y, nil
	}
}

func main() {

	if result, err := divint(5, 1); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

}
