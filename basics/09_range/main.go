package main

import "fmt"

func main() {
	ids := []int{90, 2, 21, 78, 28}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// when we do not want to use index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Range with Maps
	emails := map[string]string{"Bob": "hello@bob.com", "John": "john@gmail.com", "Sasha": "sasha@gmail.com"}

	for k, v := range emails {
		println(k + ":" + v)
	}

	for k := range emails {
		println("Name :" + k)
	}
}
