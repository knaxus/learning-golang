package main

// This is how you import multiple packages
import (
	"fmt"
	"math"

	"github.com/ashokdey/learning-golang/03_package/strutils"
)

func main() {
	fmt.Println("Hello Packages")
	size := 6.55
	fmt.Println(math.Floor(size))

	reverse := strutils.ReverseString("Hello World")

	fmt.Println(reverse)
}
