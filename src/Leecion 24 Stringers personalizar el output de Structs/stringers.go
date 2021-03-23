package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

//La estructura es func+( la variable que hereda + la estructura) + El nombre de la funcion
// + el tipo de dato que retorna
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB en disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, brand: "MSI", disk: 100}

	fmt.Println(myPC)
}
