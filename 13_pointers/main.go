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
	fmt.Println("\naddress and dereference opator:")
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	/*
	 * declaring an initializing pointers
	 * var <pointername> *<datatype> = &<variablename>
	 */
	println("\ndeclaring and initializing pointers:")
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
	fmt.Println("\ndereference pointer:")
	*mypointer = "my world"
	fmt.Printf("%T %v \n", &mystring, &mystring)

	/* passing by value
	 * pass value of var as an argument
	 * parameter is copied into another location in memory
	 * only copy is modified in function, original not
	 * basic datatype (int, bool, float, string, array) are passed by value
	 */
	fmt.Println("\npass by value:")
	a := "hello"
	fmt.Println(a)
	modifyPbV(a)
	fmt.Println(a)

	/* passing by reference
	 * pointers allow to pass references to values
	 * with call by reference/pointer address of var is passed into function call
	 * as actual parameter
	 * all operations in the function are performed on the value at the address
	 * of the actual parameter
	 */
	fmt.Println("\npass by reference:")
	modifyPbR(&a)
	fmt.Println(a)

	// slices and maps are passed be reference by default
	fmt.Println("\npass by reference is default for slices and maps:")
	slice := []int{10, 20, 30}
	fmt.Println(slice)
	modifySlice(slice)
	fmt.Println(slice)

	myMap := make(map[string]int)
	myMap["A"] = 35
	myMap["B"] = 40
	fmt.Println(myMap)
	modifyMap(myMap)
	fmt.Println(myMap)
}

// pass by value
func modifyPbV(str string) {
	str = "world"
}

// pass by reference
func modifyPbR(str *string) {
	*str = "world"
}

func modifySlice(s []int) {
	s[0] = 100
}

func modifyMap(m map[string]int) {
	m["K"] = 65
}
