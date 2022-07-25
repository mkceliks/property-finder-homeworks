package main

import "fmt"

func main() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = 2 * (width + height) // arithmetic operations to find rectangle perimeter.

	fmt.Println(perimeter)
}
