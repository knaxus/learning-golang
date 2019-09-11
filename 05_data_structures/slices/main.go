package main

import "fmt"

func main() {
	fruitSlice := []string{"Apple", "Banana", "Orange", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:2]) // will start at 1 and stop before 2
	fmt.Println(fruitSlice[1:3]) // will start at 1 and stop before 3
}
