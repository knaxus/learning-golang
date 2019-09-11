package main

import (
	"fmt"
)

func main() {
	name := "Ashok"    // when you know the type
	var age int32 = 21 // if you skip the type here, it will take int but I wanted int32

	fmt.Println(name, age)

	// using Printf to print type and we have to put new line manually
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
}
