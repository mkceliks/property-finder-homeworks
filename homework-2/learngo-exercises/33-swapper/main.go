package main

import "fmt"

func main() {
	color, color2 := "red", "blue"

	color, color2 = "orange", "green" // swapped.

	fmt.Println(color, color2)
}
