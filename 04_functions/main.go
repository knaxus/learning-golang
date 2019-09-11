package main

import (
	"fmt"
)

// when both the arguments are of same type
func add(a, b int32) int32 {
	return a + b
}

func greet(name string) string {
	return "Hello " + name + ", welcome to the island of Golang!"
}

func main() {
	fmt.Println(greet("John"))
}
