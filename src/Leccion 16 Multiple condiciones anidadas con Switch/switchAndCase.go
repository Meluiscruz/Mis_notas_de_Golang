package main

import "fmt"

func main() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	// Switch sin condición
	value := -2
	switch {
	case value > 100:
		fmt.Println("El valor es mayor a 100")
	case value < 0:
		fmt.Println("El valor es menor a 0")
	default:
		fmt.Println("No hay condición")
	}
}
