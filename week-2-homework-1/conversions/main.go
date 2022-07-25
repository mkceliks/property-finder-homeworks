package main

import "fmt"

func main() {

	var sum int = 155   // declaring sum
	var num1 int = 10   // declaring number
	var howMany float64 // declaring an empty float64 variable.

	howMany = float64(sum) / float64(num1) // type conversions to float64

	fmt.Printf("How many 10 in 155  = %f\n", howMany)
}
