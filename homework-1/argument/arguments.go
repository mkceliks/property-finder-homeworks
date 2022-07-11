package main

import (
	"fmt"
	"os" // is a library for Args.
)

func main() {

	name := os.Args[1] // Args -> when the program runs as go run arguments.go "Mustafa Kemal" with command-line.

	// "Mustafa Kemal" is a Argument here. Defining name variable is Argument.

	greeting := createGreet(name)

	fmt.Printf("%s\n", greeting)

}

func createGreet(name string) string {

	return "Merhaba " + name + " :)"

}
