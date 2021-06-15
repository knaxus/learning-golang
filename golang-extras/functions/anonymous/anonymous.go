package main

import "fmt"

// func that returns a func which returns and int
func increment(x int) func() int {
	return func() int {
		x += 1
		return x
	}
}

func main() {
	func(m string) {
		fmt.Println(m)
	}("Anonymous inside main")

	a := increment(10)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
