package main

import "fmt"

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	number := 2
	parity := isEven(number)
	fmt.Printf("Is %d an even number? : %t ", number, parity)
}
