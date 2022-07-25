package main

import "fmt" // fmt is a tool for formatting the code according to the rules of the go-lang.

type Person struct {
	name string
}

func (p Person) greet() string {
	return "Selam " + p.name
}

func main() {

	var greeter Person = Person{"Mustafa Kemal"}

	greeting := greeter.greet()
	fmt.Printf("%s\n", greeting)
}
