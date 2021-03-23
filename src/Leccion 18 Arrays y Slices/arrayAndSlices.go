package main

import "fmt"

func main() {
	//Array
	//Los arrays no son estructuras mutables
	var array [4]int
	fmt.Println(array)
	array[0] = 1
	array[1] = 2
	//fmt.Println(array)
	//fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	//fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el Slice
	//fmt.Println(slice[0])
	//fmt.Println(slice[:3])
	//fmt.Println(slice[2:4])
	//fmt.Println(slice[4:])

	//Los slice son estructuras mutables
	//Metodo append
	slice = append(slice, 7)
	fmt.Println(slice)
}
