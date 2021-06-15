package main

import "fmt"

func sum(a ...int) int {
	fmt.Printf("%T\n", a)
	return 0
}

func sumAndProduct(nums ...float64) (float64, float64) {
	sum := 0.0
	product := 1.0

	for _, v := range nums {
		sum += v
		product *= v
	}
	return sum, product
}

func increaseBy(a int, nums ...int) {
	for i := range nums {
		nums[i] *= 2
	}
}

func main() {
	// creata a slice
	nums := []int{1, 3, 5}
	nums = append(nums, 1, 9)
	decimals := []float64{1, 4, 43, 21}

	x := sum(nums...)
	y, z := sumAndProduct(decimals...)

	fmt.Println(x, y, z)

	increaseBy(2, nums...)
	fmt.Println(nums)

}
