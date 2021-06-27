package main

import "fmt"

// anonymous field names
type Book struct {
	string
	float64
	int
}

func main() {
	fmt.Println("Hello Structs - II")

	dia := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Dia", lastName: "Doe", age: 32,
	}

	fmt.Printf("%#v\n", dia)
	fmt.Printf("Dia's age is : %d\n", dia.age)

	myBook := Book{"Hello World", 20.35, 1993}
	fmt.Println(myBook)
	myBook.float64 = 45.80
	fmt.Println(myBook)
}
