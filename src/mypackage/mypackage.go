package mypackage

import "fmt"

// CarPublic con acceso público
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage es una función pública, no espera argumento alguno
func PrintMessage() {
	fmt.Println("Hola")
}

// PrintAnotherMessage espera un argumento texto, es público
func PrintAnotherMessage(text string) {
	fmt.Println(text)
}

func printPrivateMessage(text string) {
	fmt.Println(text)
}

//Nota, la convención es: Estructura pública si se nombra la primera letra
//con mauyscula.
