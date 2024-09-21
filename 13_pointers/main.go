package main

import (
	"fmt"
)

/*
* pointer is a variable that holds memory adress of another variable
* also points where memory is allocated; finding and changing values possible
 */

func main() {

	/*
	* address operator
	* & operator - '&' followed by var needed for 'adress-of operator'
	*
	* dereference operator
	* * operator - '*' followed by var returns value at the address
	 */
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	/*
	 * declaring an initializing pointers
	 * var <pointername> *<datatype> = &<variablename>
	 */

	var ptrStr *string
	fmt.Println(ptrStr)

	var ptrI *int = &i
	fmt.Println(ptrI)

	// another way
	s := "hello"
	var ptrS = &s

	fmt.Println(ptrS)

	// using shorthand operator
	mystring := "world"
	mypointer := &mystring
	fmt.Println(mypointer)
	fmt.Printf("%T %v \n", &mystring, &mystring)

	// dereferencing a pointer
	*mypointer = "my world"
	fmt.Printf("%T %v \n", &mystring, &mystring)

}
