package main

import "fmt"

func main() {

	//Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printft
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	//En caso de que no se conozca el tipo de la variable a imprimir
	fmt.Printf("%v tiene más de %d cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("cursos: %T \n", cursos)

}
