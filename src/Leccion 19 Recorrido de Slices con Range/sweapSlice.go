package main

import "fmt"

func main() {
	slice := []string{"hola", "que", "hace"}

	for _, valor := range slice {
		fmt.Println(valor)
	}
}
