package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("Ciclo %d, continue\n", i)
			continue
		}
		if i == 8 {
			fmt.Printf("Ciclo %d, break", i)
			break
		}
		println(i)
	}

}
