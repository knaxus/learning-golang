package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Car struct {
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Price float64 `json:"price"`
}

func main() {

	list := []Car{
		{Brand: "Audi", Model: "Q2", Price: 23.2},
		{Brand: "Audi", Model: "Q7", Price: 58.5},
	}

	// marshalling to JSON
	data, err := json.Marshal(list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	// unmarhslling from JSON
	var brands []struct{ Brand string }
	err = json.Unmarshal(data, &brands)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", brands)
}
