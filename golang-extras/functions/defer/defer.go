package main

import "fmt"

func foo() {
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is baar")
}

func foobar() {
	fmt.Println("This is Foobar")
}

func main() {
	fmt.Println("Using DEFER")

	defer foo()
	bar()
	defer foobar()

	// When we derfer a function, it gets paused and pushed to a STACK
	// The deferred functions will be executed in LIFO order

	fmt.Println("The last statement")
}
