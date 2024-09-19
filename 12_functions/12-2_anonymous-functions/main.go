package main

import "fmt"

/*
* anonymous function is declared without any named identifier to refer to it
* can accept inputs/outputs as standard functions
* can be used for functionalities that does not need to be named and rather for short term use
 */

func main() {
	// function is saved in a variable just a is JS/TS
	x := func(a int, b int) int {
		return a * b
	}
	fmt.Printf("%T \n", x) // return type of x
	fmt.Println(x(20, 30))
	fmt.Println("###")
	// we can also call it directly

	y := func(a int, b int) int {
		return a * b
	}(20, 30)
	fmt.Printf("%T \n", y) // return type of y; now it return int and not function becaus result is stored directly in x
	fmt.Println(x(20, 30))
}
