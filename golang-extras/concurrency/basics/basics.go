package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("F1 finished")
	wg.Done()
}

func f2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("F2 finished")
}

func main() {
	fmt.Println("Concurrency in Golang")

	fmt.Println("No. of CPUs :", runtime.NumCPU())
	fmt.Println("No. of Go routines :", runtime.NumGoroutine())

	fmt.Println("Machine OS :", runtime.GOOS)
	fmt.Println("Machine architecture :", runtime.GOARCH)

	// using go routines
	wg := sync.WaitGroup{}
	wg.Add(1)

	go f1(&wg)

	f2()
	wg.Wait()
	fmt.Println("Main finished")
}
