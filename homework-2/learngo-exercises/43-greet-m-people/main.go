package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		length  = len(os.Args) - 1
		number1 = os.Args[1]
		number2 = os.Args[2]
		number3 = os.Args[3]
	)

	fmt.Println("There are", length, "people !")
	fmt.Println("Hello great", number1, "!")
	fmt.Println("Hello great", number2, "!")
	fmt.Println("Hello great", number3, "!")
	fmt.Println("Nice to meet you all.")
}
