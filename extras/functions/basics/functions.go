package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func product(a, b int) int {
	return a * b
}

func f5(a, b int, c float64) (int, float64) {
	return a + b, c * c
}

func main() {
	a := sum(21, 31)
	b := product(10, 30)
	c, d := f5(a, b, 2.019)

	fmt.Println(c)
	fmt.Println(d)

}
