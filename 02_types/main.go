package main

import (
	"fmt"
)

func main() {
	name := "Ashok"    // when you know the type and this short hand can only be used inside a function
	var age int32 = 21 // if you skip the type here, it will take int but I wanted int32

	fmt.Println(name, age)
	// using Printf to print type and we have to put new line manually
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)

	name = "John" // can be re-assigned
	fmt.Println(name)

	const city = "Noida" // cannot assign new value to city
	fmt.Println(city)

	size := 2.4
	isTrue := false

	fmt.Println(size, isTrue)

	// save some time like:
	state, country := "UP", "India"
	fmt.Println(state, country)
}
