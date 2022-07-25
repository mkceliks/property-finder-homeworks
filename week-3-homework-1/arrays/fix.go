package main

import "fmt"

func main() {

	names := [...]string{"Einstein", "Shepard", "Tesla"} // creating a slice with infinite size ( string type )
	books := [5]string{"Kafka's Revenge", "Stay Golden"} // creating a slice with size of 5 ( string type )

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)

}
