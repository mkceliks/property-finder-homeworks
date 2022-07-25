package main

import "fmt"

func main() {

	const (
		width  = 25        // declaring
		height = width * 2 // declaring
	)

	fmt.Printf("area = %d\n", width*height)
}
