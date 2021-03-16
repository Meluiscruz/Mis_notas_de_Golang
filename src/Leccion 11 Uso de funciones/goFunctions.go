package main

import "fmt"

func normalFunction() {
	fmt.Println("Hola mundo")
}

func functionArg(message string) {
	fmt.Println(message)
}

func tripleArgs(a int, b int, c string) {
	//Es lo mismo que tripleArgs(a  b int, c string)
	fmt.Println(a, b, c)
}

func returnValue(a, b int) int {
	return a * b
}

func doubleReturn(a, b int) (c, d int) {
	return a, b * 2
}

func main() {
	normalFunction()
	functionArg("Hello World!!!")
	tripleArgs(1, 5, "Hola")

	value := returnValue(5, 4)
	fmt.Printf("The value is = %d\n", value)

	value1, value2 := doubleReturn(5, 4)
	fmt.Printf("value 1 = %d , value 2 = %d \n", value1, value2)

	value1, _ = doubleReturn(5, 4)
	//La manera de cómo escapar uno de los valores que retorna
	// una función de multiples resultados
	fmt.Printf("value 1 = %d\n", value1)
}
