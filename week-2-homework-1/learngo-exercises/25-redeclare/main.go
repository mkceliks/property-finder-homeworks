package main

import (
	"fmt"
)

func main() {
	age, yourAge := 10, 20
	age, ratio := 42, 3.14 // redeclaration of age.

	fmt.Println(age, yourAge, ratio)
}
