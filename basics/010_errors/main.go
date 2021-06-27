package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func getGreeting(hours int) (string, error) {
	var message string

	if hours < 7 {
		err := errors.New("Too earlly for greetings")
		return message, err
	}

	if hours < 12 {
		message = "Good Morning!"
	} else if hours < 18 {
		message = "Good Afternoon!"
	} else {
		message = "Good Evening"
	}
	return message, nil
}

func main() {
	hourOfTheDay := time.Now().Hour()
	greetings, err := getGreeting(hourOfTheDay)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(greetings)
}
