package main

import (
	"fmt"
	"strings"
)

func main() {

	greetPrinter(createGreetInTurkish, "greeting in turkish") // We are using function inside a function.
	greetPrinter(createGreetInEnglish, "greeting in english")
	greetPrinter(convertToUpperCase, "upper case made")

	func(name string) { // this is a local anonymus function that we can use in that time.
		greeting := "Merhaba " + name
		fmt.Printf("%s\n", greeting)
	}("Mustafa Kemal") // this is a argument "name string".

	closure := func(name string) { // named anonymus function with "closure".
		greeting := "Merhaba " + name
		fmt.Printf("%s\n", greeting)
	}
	closure("Kemal") // calling the anonymus function with named "closure".

	anotherGreetPrinter(closure, "Mustafa Kemal with closure")
}

func createGreetInTurkish(name string) string {
	return "Merhaba " + name
}

func createGreetInEnglish(name string) string {
	return "Hello " + name
}

func convertToUpperCase(arg string) string {
	return strings.ToUpper(arg)
}

func greetPrinter(function func(it string) string, name string) { // function recieving a string and returns a string.
	var greeting = function(name)
	fmt.Printf("%s\n", greeting)
}

func anotherGreetPrinter(function func(it string), name string) {
	function(name)
}
