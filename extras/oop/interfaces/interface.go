package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length, breadth float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * r.length * r.breadth
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func print(s Shape) {
	fmt.Printf("Shape : %#v\n", s)
	fmt.Printf("Area : %v\n", s.Area())
	fmt.Printf("Perimeter : %v\n", s.Perimeter())
}

func main() {

	var s Shape
	// what's the type of s?
	fmt.Printf("%T\n", s)

	s = Rectangle{105.45, 34.65}
	s.Area()

	s = Circle{10.05}
	s.Perimeter()

	print(s)

	// type assertion
	s = Circle{20.35}

	// we cannot call s.volume() because the interface
	// Shape has no method volume()
	// Using type assertion & switch

	s.(Circle).volume()

	// the above can fail based on what's coming dynamically
	ball, ok := s.(Circle)
	var volume float64

	if ok {
		volume = ball.volume()
		fmt.Printf("Volume of ball : %v\n", volume)
	}

	// other way to handle the same is using a switch case
	switch value := s.(type) {
	case Circle:
		fmt.Printf("%#v has a circle\n", value)
	case Rectangle:
		fmt.Printf("%#v has a circle\n", value)
	}

}
