package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 1988
	fmt.Println(myCar)

	//var myOtherCar pk.carPrivate
	//fmt.Println(myOtherCar)

	pk.PrintMessage()
	pk.PrintAnotherMessage("Esto es un argumento")

	//pk.printPrivateMessage("Esto es secreto")

}
