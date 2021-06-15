package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var totalIterations int32 = 0
var m sync.Mutex

// initialization
func init() {
	cpuAvailable := runtime.NumCPU()
	cpuToUse := 4
	runtime.GOMAXPROCS(cpuToUse)

	fmt.Printf("Total CPU : %d | Using CPU: %d\n", cpuAvailable, cpuToUse)
}

func number(id string, wg *sync.WaitGroup) {
	// seed  is important to create a random seed value
	rand.Seed(time.Now().Unix())

	// print random nubers
	for i := 0; i < 10; i++ {
		m.Lock()
		totalIterations++
		m.Unlock()

		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Thread %s : %d \n", id, rand.Intn(20)+20)
	}
	wg.Done()
}

func alphabets(wg *sync.WaitGroup) {
	for i := 'A'; i <= 'J'; i++ {
		m.Lock()
		totalIterations++
		defer m.Unlock()

		time.Sleep(400 * time.Millisecond)
		fmt.Printf("Character : %c \n", i)
	}
	wg.Done()
}

func main() {
	startTime := time.Now()

	ids := []string{"1", "2", "3", "4"}

	wg := sync.WaitGroup{}
	wg.Add(len(ids) + 1)

	for i := range ids {
		go number(ids[i], &wg)
	}
	go alphabets(&wg)

	wg.Wait()

	fmt.Println("Total iterations : ", totalIterations)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total time: %s\n", elapsedTime)
}
