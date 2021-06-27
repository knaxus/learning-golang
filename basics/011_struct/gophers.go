package main

import (
	"fmt"
)

type gopher struct {
	name string
	age  int
}

// Adding methods to Gopher

func (g gopher) jump() string {
	if g.age < 45 {
		return g.name + " can jump HIGH!"
	}

	if g.age < 90 {
		return g.name + " can still jump..."
	}

	return g.name + " cannot jump..."
}

func main() {
	gopherOne := gopher{name: "Sam", age: 15}
	gopherTwo := gopher{name: "John", age: 65}
	elderGopher := gopher{name: "Phil", age: 95}

	fmt.Println(gopherOne.jump())
	fmt.Println(gopherTwo.jump())
	fmt.Println(elderGopher.jump())
}
