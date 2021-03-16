package main

import "fmt"

func validateUsername(username, password string) string {
	if username == "username" && password == "password" {
		return "Succsess!!!!\n"
	} else {
		return "Your username or pasword is invalid!\n"
	}
}

func main() {
	username := "userMame"
	password := "password"
	fmt.Printf(validateUsername(username, password))

}
