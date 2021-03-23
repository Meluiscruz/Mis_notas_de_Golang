package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pedro"] = 20

	fmt.Println(m)

	//Recorrer map
	for key, v := range m {
		fmt.Println(key, v)
	}

	//Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
