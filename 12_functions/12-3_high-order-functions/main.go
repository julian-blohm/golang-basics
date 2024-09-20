package main

import (
	"fmt"
	"math"
)

/*
* receives a function as argument or returns function as output
* why such functions?
* - composition: create smaller function for certain piece of logic; compose "complexity" with smaller functions
* - reduce bugs (more modular)
* - code gets easier to read and understand
 */

// example: calculating some things for a circle

func calcArea(r float64) float64 {
	return math.Pi * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * math.Pi * r
}

func calcDiameter(r float64) float64 {
	return 2 * r
}

// high order function with radius and function as input
func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you!")
}

// int parameter and function as return
func getFunction(query int) func(r float64) float64 {
	// maps ints to functions
	queryToFunc := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return queryToFunc[query]
}

func main() {
	var query int // to define what to calculate
	var radius float64

	fmt.Println("Enter the Radius of the circle: ")
	fmt.Scanf("%f", &radius) // get float input with %f fomatter
	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter \n")
	fmt.Scanf("%d", &query) // get int input with %d fomatter

	printResult(radius, getFunction(query))
}
