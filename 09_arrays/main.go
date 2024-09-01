package main

import (
	"fmt"
)

// array is a collection of similar data elements stores at contiguous memory location e.g. ints, stirngs...
// arrays have a fixed length; elements should be of same data type

// declare array:
// var <array name>[<size of array>] <data type>
var grades [5]int
var fruits [3]string

func main() {
	fmt.Println(grades) // only contains 0 because int and not initialized
	fmt.Println(fruits) // contains nothing because string and not initialized

	// to initialize array; length, datatype and values (less or egauel size) needed

	grades = [5]int{1, 2, 3, 4, 5}
	fruits = [3]string{"apple", "banana", "ananas"}

	// also possible that compiler defines size of array based on number of elements; [...] is called ellipsis

	ages := [...]int{15, 30, 41}

	fmt.Println(grades)
	fmt.Println(fruits)
	fmt.Println(ages)

	// find length of array uing len() method
	fmt.Println("Size of fruit array = ", len(fruits))

	// arrays and indexes
	fmt.Println("Second fruit in array (with index 1) is: ", fruits[1])

	// overwrite value in array
	fruits[1] = "mango"
	fmt.Println("Element on with index 1 is now: ", fruits[1])

	// loop over array
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	// looping through array with range keyword
	for index, element := range grades {
		counter := 1
		fmt.Println("Iteration numner", counter)
		fmt.Println(index, "=>", element)
		counter += 1
	}

	// multidimensional arrays (array of arrays => 2d array)
	arr := [3][2]int{
		{2, 4},
		{4, 16},
		{8, 64},
	}
	fmt.Println(arr[2][1])

}
