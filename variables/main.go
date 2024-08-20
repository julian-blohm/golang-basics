package main

import (
	"fmt"
)

var hi string = "hi" // global variable

func main() {
	// these variables are only accessible within this block surrounded by {}
	// it is still possible to access variables from outer blockes
	var helloWorld string = "Hello World"
	var myIntNumber int = 500
	var myBool bool = false     // local variable if inside function or lood... declared
	var myFloat float64 = 74.87 // best practice to use float64 instead of float32

	// not in new line, both will be in same line
	fmt.Print(hi)
	fmt.Print(myIntNumber)

	// following possible
	fmt.Print(myFloat, "\n")
	fmt.Print(myBool, "\n")

	// but this is better
	fmt.Println(helloWorld)
	fmt.Println("My Int Number is: ", myIntNumber)
	fmt.Println(myBool, " is my boolean and ", myFloat, " is my float")

	// to format variables for output with printf
	/*
	 * format specifier
	 * %v - formats value in default format
	 * %d - format decimal integers
	 * find more here about the fmt package and printf https://pkg.go.dev/fmt
	 */
	fmt.Printf("Hi %v", helloWorld)
	fmt.Printf("MyIntNumber: %d", myIntNumber)
	fmt.Println("")

	// use together
	var age int = 42
	var name string = "bob"
	fmt.Printf("My name is %v and I am %d years old", name, age)

	/*
	 * btw declaration and assigning value can be split as in java or c
	 * it is important that incorrect values for specific datatypes throw errors
	 * myNewName = 123 would throw an error
	 */
	var myNewName string
	myNewName = "joe"
	println(myNewName)

	// there is also a shorthand way to declare multiple values of same data type
	var name1, name2 string = "name1", "name2"
	println(name1, name2)

	// if different datatype
	var (
		name3 string = "name3"
		age3  int    = 80
	)
	println(name3, " ", age3)

	/*
	 * there is another way: shorthand variable declaration
	 * no datatype declaration is needed here
	 * the datatype will be assigned automatically by the compiler
	 */
	name4 := "name4"
	name4 = "newname4" // reassignement possible with same datatype, not int here possible
	println(name4)
	fmt.Println("")

	/*
	 * zero value variables
	 * bool => false
	 * int => 0
	 * float64 => 0.0
	 * string => ""
	 * pointers, functions, maps, interfaces => nil
	 */
	var zeroValueInt int
	var zeroValueFloat float64
	fmt.Printf("%d \n", zeroValueInt)
	fmt.Printf("%.f \n", zeroValueFloat)

}
