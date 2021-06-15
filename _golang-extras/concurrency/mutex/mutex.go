package main

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

	// get a mutex to lock `n` and allow single acess at a time
	m := sync.Mutex{}

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock() // same as the above function
			n--
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of n ->", n)
}
