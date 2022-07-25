package main

import (
	"fmt"
)

type Person struct { // defining a struct.
	name string // data inside of struct.
}

func (p Person) greet() string { // function greet() is associated with struct Person and thus becomes a method.
	return "Merhaba " + p.name
}

func main() {

	greeter := Person{"Mustafa Kemal"} // defining a data with the type of Person struct
	var greeting = greeter.greet()
	fmt.Printf("%s\n", greeting)

}
