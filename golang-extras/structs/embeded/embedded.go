package main

import "fmt"

type Name struct {
	first, middle, last string
}

type Contact struct {
	email, address string
	pincode        int
}

type Employee struct {
	fullName    Name
	designation string
	contact     Contact
}

func main() {
	fmt.Println("Hello Structs E:3")

	john := Employee{
		fullName: Name{
			first: "John",
			last:  "Doe",
		},
		designation: "Manager",
		contact: Contact{
			email:   "hello@doe.com",
			address: "4/3 Cross Street",
			pincode: 210901,
		},
	}

	fmt.Printf("Enployee is %s \nEmail %s\n", john.fullName.first, john.contact.email)
}
