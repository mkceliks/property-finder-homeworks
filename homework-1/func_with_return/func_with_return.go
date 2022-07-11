package main

import (
	"fmt"
)

func main() {

	var name string = "Mustafa Kemal" // parameter is "name" variable.

	var greeting = createGreet(name) // greeting variable is declaring like return value of createGreet function.

	fmt.Printf("%s", greeting) // printing the greeting variable. ( Return of createGreet function. )

}
func createGreet(name string) string {

	greeting := "Merhaba " + name + " Mustafa Kemal" // declaring a variable with a name "greeting". ":= is for declaration with assignment, whereas = is for assignment only."

	return greeting // greeting variable is a return value here.

}
