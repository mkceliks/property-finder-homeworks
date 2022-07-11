package main

import (
	"fmt"
)

func main() {
	var num1 int = 13

	var num2 int = 7

	var num3 int

	num4 := 5

	fmt.Println(num1) // 13
	fmt.Println(num2) // 7
	fmt.Println(num3) // 0
	fmt.Println(num4) // 5

	num3 = num1 + num2 // 13 + 7
	fmt.Println(num3)  // num3 s= 20

	num3 = num1 - num2 // 13-7
	fmt.Println(num3)  // num3 = 6

	num3 = num1 % num4 // 13 mod 5
	fmt.Println(num3)  // 3

}
