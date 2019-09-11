package main

/*
1. Has to have fixed length
2. Has to have a type
*/

import (
	"fmt"
)

func main() {
	var fruits [2]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"

	myFav := fruits[1]

	fmt.Println(fruits)
	fmt.Println(myFav)

	// Declare and assign together
	cities := [3]string{"Noida", "Delhi", "Gurugram"}
	fmt.Println(cities)
}
