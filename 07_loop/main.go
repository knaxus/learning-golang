package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang")

	// In golang, there is only one loop, the For Loop

	for i := 0; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		}
	}

	// Replace while loop
	z := 3
	for z <= 5 {
		fmt.Println(z)
		z++
	}
}
