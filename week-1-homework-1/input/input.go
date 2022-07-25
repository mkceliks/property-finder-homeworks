package main

import (
	"fmt"
)

func main() {

	var name string

	fmt.Println("Please enter your name:")

	fmt.Scanln(&name) // An input can also be received from the user using Scanln function of fmt library.

	greeting := createGreet(name)

	fmt.Printf("%s\n", greeting)

}
func createGreet(name string) string {

	return "Merhaba " + name

}
