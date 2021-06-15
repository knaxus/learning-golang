package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filename := "main.go"

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)

	fmt.Println("Files in Golang")
}
