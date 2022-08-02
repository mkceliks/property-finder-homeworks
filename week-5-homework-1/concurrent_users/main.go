package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Usage interface { // Interface for the usage of the system
	Limit()
}

type Licence struct{} //

type User struct { // User of the system
	Age  int
	Name string
}

type LicenceProxy struct { // Proxy for the user
	licence Licence
	user    *User
}

func (c *Licence) Limit(wg *sync.WaitGroup, number int) { // Limit the number of concurrent users
	fmt.Printf("Number of concurrent users %d. All of them using the system very well.\n", number)
	wg.Done()
}

func NewLimitProxy(user *User) *LicenceProxy { // Create a proxy for the user
	return &LicenceProxy{Licence{}, user}
}

func (c *LicenceProxy) Limit(wg *sync.WaitGroup, number int) { // Limit the number of concurrent users
	if number <= 5 { // If the number of concurrent users is less than 5
		c.licence.Limit(wg, number)
	} else {
		fmt.Printf("Number of concurrent users %d. Please try again later.\n These people cannot reach the system right now : %s.\n", 5, c.user.Name) // If the number of concurrent users is greater than 5, print the message
		wg.Done()
	}
}

func main() {

	users := []User{ // Users of the system
		{
			Age:  10,     // Age of the user
			Name: "John", // Name of the user
		},
		{
			Age:  20,
			Name: "Jane",
		},

		{
			Age:  30,
			Name: "Jack",
		},
		{
			Age:  40,
			Name: "Jill",
		},
		{
			Age:  50,
			Name: "Joe",
		},
		{
			Age:  60,
			Name: "Juan",
		},
		{
			Age:  70,
			Name: "Josh",
		},
		{
			Age:  80,
			Name: "Jacky",
		},
	}

	for i := 0; i < len(users); i++ { // For each user
		licence := NewLimitProxy(&users[i]) // Create a proxy for the user
		wg.Add(1)                           // Add one to the wait group
		go licence.Limit(&wg, i+1)          // Limit the number of concurrent users
		wg.Wait()                           // Wait for the user to finish
	}

}
