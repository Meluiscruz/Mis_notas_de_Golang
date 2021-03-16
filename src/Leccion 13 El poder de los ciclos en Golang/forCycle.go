package main

import "fmt"

func main() {

	//Ciclo For condicional exclusivo
	//inclusivo: i <= 10
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	//Ciclo For While
	counter := 0
	for counter < 14 {
		fmt.Println(counter)
		counter++
	}

	fmt.Printf("\n")

	//Ciclo Forforever, se detiene con Ctrl + C
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
