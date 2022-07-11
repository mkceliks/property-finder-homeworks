package main

import "fmt"

func main() {

	a, b := 10, 5.5

	a = int(b) // a is 5 here

	fmt.Println(float64(a) + b) // a converted to float64 from int and addition successfull.

}
