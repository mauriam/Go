package main

import "fmt"

func main() {
	//bucle infinito
	/*
		for {
			println("Hola mundo")
		}
	*/
	// bucle tipo while
	numero := 124558564512
	c := 0
	for numero > 0 {
		numero /= 10
		c++
	}
	fmt.Println("La cantidad de digitos es: ", c)
	// bucle tipo for
	for i := 0; i < 10; i++ {
		fmt.Println("ciclo:", i)
	}
	// bucle tipo for con if
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("ciclo par:", i)
		}
	}

}
