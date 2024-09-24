package main

import "fmt"

/*
 * interface in go specifies a method set and is especially for better modularity
 * blueprint for method set
 * they specify a set of methods but do not implement them
 * a type implements an interface by implementing the methods
 * go interfaces are implemented implicitly
 * it does not need any specific keyword to implement an interface
 */

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (sq square) area() float64 {
	return sq.side * sq.side
}

func (sq square) perimeter() float64 {
	return 4 * sq.side
}

type rect struct {
	length, breadth float64
}

func (re rect) area() float64 {
	return re.length * re.breadth
}

func (re rect) perimeter() float64 {
	return 2*re.length + 2*re.breadth
}

func printData(sh shape) {
	fmt.Println(sh)
	fmt.Println(sh.area())
	fmt.Println(sh.perimeter())
}

func main() {
	r := rect{length: 3, breadth: 4}
	s := square{side: 5}
	printData(r)
	printData(s)
}
