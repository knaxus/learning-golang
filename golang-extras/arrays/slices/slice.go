package main

import (
	"fmt"
)

type Any struct{}

func main() {
	// make a slice
	sl := make([]Any, 5, 10)

	fmt.Println(sl)

}
