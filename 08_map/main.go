package main

import "fmt"

func main() {
	// Define map
	cityStates := map[string]string{"Noida": "Uttar Pradesh", "Pune": "Maharastra"}
	fmt.Println(cityStates)

	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["John"] = "john@gmail.com"

	fmt.Println(emails)         // get all the key-value
	fmt.Println(emails["John"]) // get single value
	fmt.Println(len(emails))    // get length

	emails["Dan"] = "danny@don.co"
	fmt.Println(emails)

	// delete entry from Map
	delete(emails, "John")

	fmt.Println(emails)
}
