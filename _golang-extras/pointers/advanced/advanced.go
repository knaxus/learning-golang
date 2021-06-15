package main

import "fmt"

type Product struct {
	name     string
	quantity int
}

func changeProduct(p *Product) {
	p.name = "Random"
	p.quantity = 101
}

func main() {
	fmt.Println("Advanced Pointers")

	gift := Product{name: "Candles", quantity: 10}
	fmt.Println(gift)

	// change the values
	changeProduct(&gift)

	fmt.Println(gift)
}
