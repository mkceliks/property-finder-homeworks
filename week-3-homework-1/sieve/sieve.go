package main

import (
	"fmt"
)

func display(num int, prime []bool) { // display function which takes 2 arguments.
	fmt.Printf("Primes less than %d : ", num)
	for i := 2; i <= num; i++ {
		if prime[i] == false {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func sieve(num int) { // Sieve of Eratosthenes Algorithm implementation with given number.
	prime := make([]bool, num+1)
	for i := 0; i < num+1; i++ {
		prime[i] = false
	}
	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true
			}
		}
	}
	display(num, prime)
}

func main() {
	var n int
	fmt.Println("Enter a number for Sieve of Eratosthenes Algorithm : ")
	fmt.Scanf("%d", &n) // taking an integer from user.
	sieve(n)
}
