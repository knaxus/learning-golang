package main

import (
	"fmt"
	"math"
)

type Solid interface {
	Volume() float64
}

type TwoDimensional interface {
	Area() float64
}

// embedded interface geometry
type Geometry interface {
	Solid
	TwoDimensional
	getColor() string
}

type Cube struct {
	side  float64
	color string
}

func (c Cube) Area() float64 {
	return 6 * math.Pow(c.side, 2)
}

func (c Cube) Volume() float64 {
	return math.Pow(c.side, 3)
}

func (c Cube) getColor() string {
	return c.color
}

func measurements(g Geometry) (float64, float64) {
	return g.Area(), g.Volume()
}

func main() {
	fmt.Println("Embedded interaces in Golang")

	cube := Cube{35, "red"}

	a, v := measurements(cube)

	fmt.Printf("Area : %v, Volume : %v\n", a, v)

}
