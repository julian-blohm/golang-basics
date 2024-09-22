package main

import (
	"fmt"
	"math"
)

/*
 * method augments function and adds extra parameter section after func keyword
 * func (<receiver>) <methodname>(<parameters>) <returnparams> { <code> }
 * receiver can use a type oder also pointer ot a named type
 */

type Circle struct {
	radius float64
	area   float64
}

// following functions are the method set of the struct
func (c Circle) calcArea() {
	c.area = math.Pi * c.radius * c.radius
}

func (c *Circle) calcAreaPTR() {
	c.area = math.Pi * c.radius * c.radius

}

func (c Circle) printRadius() {
	fmt.Println(c.radius)
}

func (c Circle) printArea() {
	fmt.Println(c.area)
}

func main() {
	c := Circle{radius: 5}
	c.calcArea()
	fmt.Printf("%+v", c)

	fmt.Println("")

	c.calcAreaPTR()
	fmt.Printf("%+v", c)
}
