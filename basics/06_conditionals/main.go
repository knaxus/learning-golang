package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to conditionals in Go!")

	x, y, z := 5, 10, 8

	if x < y {
		fmt.Println("Oh yeah!")
	}

	if x > y {
		fmt.Println("Oh no..")
	} else if z > x {
		fmt.Println("This is nice!")
	} else {
		fmt.Println("Gone.")
	}

	color := "blue"

	switch color {
	case "blue":
		fmt.Println("Captain America!")
	case "red":
		fmt.Println("Iron Man")
	case "green":
		fmt.Println("Hulk")
	}

	if z > x && color == "blue" {
		fmt.Println("Great, move on!")
	}

}
