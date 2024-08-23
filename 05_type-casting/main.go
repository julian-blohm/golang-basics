package main

import (
	"fmt"
	"strconv"
)

func main() {

	// integer to float
	var intNumber = 90
	var floatNumber = float64(intNumber)
	fmt.Printf("%.2f\n", floatNumber)

	// float to integer
	floatNumber = 45.89
	intNumber = int(floatNumber)
	fmt.Printf("%v", intNumber)

	// string conversion with package strconv
	// itoa() method: converts integer to string
	var myString = strconv.Itoa(intNumber)
	fmt.Printf("%q", myString)

	// atoi() method: converts string to integer
	// not valid values will return an error
	myString = "5"
	intNumber, err := strconv.Atoi(myString)
	fmt.Printf("%v, %T\n", intNumber, intNumber)
	fmt.Printf("%v, %T\n", err, err)
}
