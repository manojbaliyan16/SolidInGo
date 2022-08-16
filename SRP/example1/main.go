package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() {
	// violating the SRP due to below reasons
	// 1. Out may need to be represent into more than one format
	// 2. Calculation of area may need to be change.
	fmt.Printf("The area of the circle is %f", math.Pi*c.radius*c.radius)

}

func main() {
	c := Circle{
		radius: 3,
	}
	c.area()
}
