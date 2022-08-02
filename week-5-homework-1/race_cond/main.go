package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0 // counter is a variable that is shared by all goroutines

	const num = 1000      // num is a constant that is shared by all goroutines
	var wg sync.WaitGroup // wg is a WaitGroup that is shared by all goroutines

	for i := 0; i < num; i++ {
		wg.Add(2)   // Add two to the WaitGroup
		go func() { // anonymous function that is executed by a goroutine
			temp := counter
			temp++ // increment the counter
			counter = temp
			wg.Done() // signal that this goroutine is done
		}()
		go func() { // anonymous function that is executed by a goroutine
			temp := counter
			temp-- // decrement the counter
			counter = temp
			wg.Done() // signal that this goroutine is done
		}()
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Println("count:", counter)
}
