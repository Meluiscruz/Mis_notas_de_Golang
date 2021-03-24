package main

import "fmt"

func say(text string, c chan<- string) { //Este canal solo es de entrada de datos chan<-
	c <- text
}

func main() {
	c := make(chan string, 1) //Este canal solo va a recibir una goroutine a la vez

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)
}
