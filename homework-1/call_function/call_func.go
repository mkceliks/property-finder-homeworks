package main

import (
	"fmt"
)

func main() { // main function can call other functions such as greet. Main function takes no parameters and returns no value.

	greet("Mustafa Kemal") // "Mustafa Kemal" is a parameter here. We calling greet function with a parameter "Mustafa Kemal".

}

func greet(name string) { // This is a function of outside of main function.

	fmt.Printf("Merhaba benim adım : %s\n", name)

}
