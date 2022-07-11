package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 150                                   // declaring a number
	str := strconv.Itoa(num)                        // changing the type of number with strconv.
	fmt.Printf("Type : %T\nValue : %v\n", str, str) // Type: String Value :150
}
