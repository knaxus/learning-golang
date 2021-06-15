package main

import "fmt"

func sendAndReceive(n chan int) {
	for i := 0; i < 10; i++ {
		n <- i + 1
	}
}

func main() {
	fmt.Println("Learning channel")

	num := make(chan int)
	defer close(num)

	go sendAndReceive(num)

	for i := 0; i < 20; i++ {
		value, open := <-num
		if !open {
			break
		}
		fmt.Println(value)
	}
}
