package main

import (
	"fmt"
)

func main() {
	const (
		counter1 = iota // 0
		counter2
		counter3
		counter4
		counter5
		counter6
	)
	fmt.Println(counter1, counter2, counter3, counter4, counter5, counter6) // 0,1,2,3,4,5
}
