package main

import (
	"fmt"
)

func main() {

	/*
	 * get user input of type string using the %s format specifier for string
	 * why the name variable has a & before it will be focused on later, basically it is about pointers
	 */

	var name string
	fmt.Print("Enter your name: ")

	fmt.Scanf("%s", &name)

	fmt.Println("Hey there, ", name)

	// handling multiple inputs
	var newName string
	var age int

	fmt.Print("Enter your new name & how old are you? ")

	// input is read seuqentially
	fmt.Scanf("%s %d", &newName, &age)
	fmt.Println(newName, age)

	/*
	* Scanf return values
	* count: number of arguments
	* err: any error thrown during execution
	 */

	var a string
	var b int
	fmt.Print("Enter a string and a number: ")

	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count: ", count)
	fmt.Println("error: ", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
