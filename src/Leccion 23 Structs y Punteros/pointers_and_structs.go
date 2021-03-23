package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {

	fmt.Println(myPC.brand, "pong")
}

func (myPC *pc) duplicateRAM() {

	myPC.ram = myPC.ram * 2

}

func main() {

	myPC := pc{ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
}
