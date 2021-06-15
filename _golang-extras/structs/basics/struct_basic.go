package main

import "fmt"

// define a book
type Book struct {
	title  string
	author string
	year   int
}

func main() {
	fmt.Println("Hello Structs")

	myBook := Book{"SICP", "Random", 1998}
	bestBook := Book{title: "Algorithms", author: "Thomac C", year: 1995}

	fmt.Println(myBook)
	fmt.Println(bestBook)

	// struct i copy by value
	randomBook := myBook

	myBook.title = "Database internals"

	fmt.Printf("%+v\n", randomBook)
	fmt.Printf("%+v\n", myBook)
}
