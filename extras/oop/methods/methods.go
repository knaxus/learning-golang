package main

import "fmt"

// create a type
type Names []string

// method for type names
func (n Names) print() {
	for i, name := range n {
		fmt.Printf("%d : %s\n", i, name)
	}
}

type Car struct {
	brand string
	price float64
}

func (c *Car) changeCar(brand string, price float64) {
	c.brand = brand
	c.price = price
}

func main() {
	fmt.Println("Method receivers in Golang")

	friends := Names{"Tonny", "Thor"}

	// using as a method
	friends.print()
	// passing as argument
	Names.print(friends)

	// now we need to change the properties of the object
	car := Car{"BMW", 101000}
	fmt.Println(car)

	car.changeCar("Audi", 98000)
	fmt.Println(car)
}
