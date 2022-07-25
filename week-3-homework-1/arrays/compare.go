package main

import "fmt"

func main() {
	week := [...]string{"Monday", "Tuesday"}  // creating a slice with an infinite size
	wend := [...]string{"Saturday", "Sunday"} // creating a slice with an infinite size

	fmt.Println(week != wend)

	evens := [...]int{2, 4, 6, 8, 10}  // creating a slice with an infinite size
	evens2 := [...]int{2, 4, 6, 8, 10} // creating a slice with an infinite size

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]uint8{'h', 'i'} // creating a slice with a size of 5 ( unsigned integer-8- type. )
	var data [5]byte

	fmt.Println(data == image)
}
