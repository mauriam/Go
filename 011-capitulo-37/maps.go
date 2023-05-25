package main

import "fmt"

func main() {
	//crear un mapa de estudiantes
	estudiantes := make(map[string][]int)
	//agregar estudiantes y sus notas
	estudiantes["Eduardo"] = []int{10, 8, 9, 7, 10}
	estudiantes["Juan"] = []int{6, 5, 4, 7, 8}
	estudiantes["Pedro"] = []int{10, 10, 9, 10, 10}
	//recorrer el mapa
	fmt.Println(estudiantes)
	//obtener las notas de un estudiante por su clave
	fmt.Println("Eduardo", estudiantes["Eduardo"])
	//obtener las notas de un estudiante por su clave y posición
	fmt.Println("Eduardo", estudiantes["Eduardo"][0])

}
