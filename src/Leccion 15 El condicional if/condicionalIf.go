package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("El valor es 1")
	} else {
		fmt.Println("El valor es distinto a 1")
	}

	//AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("AND: Esto es verdadero")
	}

	//OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("OR: Esto es verdadero")
	}
}
