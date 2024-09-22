package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

func calcArea(c Circle) {
	// const PI float64 = 3.14
	var area float64
	area = math.Pi * c.radius * c.radius
	c.area = area

}

func calcAreaPTR(c *Circle) {
	// const PI float64 = 3.14
	var area float64
	area = math.Pi * c.radius * c.radius
	(*c).area = area

}

func main() {
	c := Circle{x: 5, y: 5, radius: 5, area: 0}

	fmt.Println("calcArea Function")
	fmt.Printf("%+v \n", c)
	calcArea(c)
	fmt.Printf("%+v \n", c)

	fmt.Println("calcAreaPTR Function")
	fmt.Printf("%+v \n", c)
	calcAreaPTR(&c)
	fmt.Printf("%+v \n", c)

}
