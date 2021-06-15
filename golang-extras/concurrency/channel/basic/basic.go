package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	// create a channel (bidirectional by default)
	c := make(chan int)

	// create a receiver channel
	cr := make(<-chan string)

	// create a sender channel
	cs := make(chan<- string)

	fmt.Printf("%T , %T, %T \n", c, cr, cs)

	// send a value to the channel
	go f1(10, c)

	// receive the value from the channel
	n := <-c

	fmt.Println("Received from channel: ", n)
	fmt.Println("Exiting...")
	close(c)
	close(cs)
}
