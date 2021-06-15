package main

import "fmt"

func main() {
	name := "John Doe"
	fmt.Println(&name)

	p := &name
	fmt.Println(*p)

	// pointer to a pointer
	fmt.Println("Pointer to a pointer")

	a := 5.5
	p1 := &a

	p2 := &p1

	fmt.Println(*p1)
	fmt.Println(**p2)

}
