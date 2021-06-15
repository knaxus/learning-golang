package main

// execute this with the race flag
// go run -race  data_race/race.go

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	gr := 100
	n := 0

	wg := sync.WaitGroup{}
	wg.Add(gr * 2)

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of n ->", n)
}
