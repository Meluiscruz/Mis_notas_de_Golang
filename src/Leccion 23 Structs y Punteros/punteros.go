package main

import "fmt"

func main() {
	a := 50
	b := &a //Acceder a la memoria (dirección)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) //Acceder al valor almacenado en la memoria

	*b = 100
	fmt.Println(a)

}
