package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down!\n", url)
	} else {
		defer res.Body.Close()
		fmt.Printf("%s -> %d\n", url, res.StatusCode)
		if res.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}
			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing response body to : %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {

	urls := []string{"https://google.com", "https://golang.org", "https://medium.com"}
	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
