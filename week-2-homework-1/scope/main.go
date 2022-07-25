package main

import (
	"fmt"
)

var name string // global variable.

func printer() {

	name = "Mustafa Kemal Çelik" // local variable.
}

func main() {

	fmt.Println(name) // Empty

	printer() // name initialized

	fmt.Println(name) // Mustafa Kemal Çelik
}
