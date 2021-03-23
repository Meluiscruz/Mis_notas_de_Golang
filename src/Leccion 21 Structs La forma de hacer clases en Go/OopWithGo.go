package main

import "fmt"

//Los structs se crean afuera de la función main()
type car struct {
	//Aqui van los atributos
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020} //Una manera de instanciar
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
