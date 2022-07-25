package main

import (
	"fmt"
	"time"
)

func main() {

	var pizza []string

	// Departure times
	var departures []time.Time

	// Student graduation years
	var grads []int

	// On/off states of lights in a room
	var lights []bool

	pizza = append(pizza, "pepperoni", "onions", "extra cheese")

	now := time.Now()
	departures = append(departures,
		now,
		now.Add(time.Hour*24), // 24 hours after `now`
		now.Add(time.Hour*48)) // 48 hours after `now`

	grads = append(grads, 1998, 2005, 2018)

	lights = append(lights, true, false, true)

	fmt.Printf("pizza       : %s\n", pizza)
	fmt.Printf("\ngraduations : %d\n", grads)
	fmt.Printf("\ndepartures  : %s\n", departures)
	fmt.Printf("\nlights      : %t\n", lights)
}
