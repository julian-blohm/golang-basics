package main

import (
	"fmt"
	"reflect"
)

func main() {
	var grades int = 42
	var message string = "hello world"
	var isCheck bool = true
	var amount float32 = 5644.54

	// types can be checked by using the %T format specifier
	// %v is used to get the value of the var
	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = '%v' is of type %T \n", message, message)
	fmt.Printf("variable isCheck = '%v' is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)

	// also possible to use the reflect package and the TypeOf() method to get datatype of values
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Hello World"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(3.14))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))

	// can also be used for variables
	fmt.Printf("variable grades = %v is of type %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable message = '%v' is of type %v \n", message, reflect.TypeOf(message))
	fmt.Printf("variable isCheck = '%v' is of type %v \n", isCheck, reflect.TypeOf(isCheck))
	fmt.Printf("variable amount = %v is of type %v \n", amount, reflect.TypeOf(amount))
}
