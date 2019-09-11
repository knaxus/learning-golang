package main

// This is how you import multiple packages
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello Packages")
	size := 6.55
	fmt.Println(math.Floor(size))
}
