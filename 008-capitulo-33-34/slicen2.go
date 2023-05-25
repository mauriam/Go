package main

import "fmt"

func main() {
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio"}
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	//agregar elemento al slicen
	meses = append(meses, "Julio")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	//agregar elemento al slicen
	meses = append(meses, "Agosto")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
}
