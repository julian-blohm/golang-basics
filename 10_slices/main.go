package main

import (
	"fmt"
)

// slice is a contiguous sequence of an underlying array
// variable typed, more flexible

func main() {
	// declaration and initialization of slice is similar to array; but size mandatory
	grades := []int{1, 2, 3, 4, 5}
	fmt.Println(grades)

	// usually slices are created based on an array defining start of the index and the end of it
	// lets create an array

	myArray := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// create slice from it; include elements from index 1 to index 7 (upper boundary always not included)
	mySlice := myArray[1:8]
	fmt.Println(mySlice)

	// create slice from slice with element index 0 to index 2
	subSlice := mySlice[0:3]
	fmt.Println(subSlice)

	// declaring slice with make function; following parameter: datatype, length, capacity (optional)
	newSlice := make([]int, 5, 10)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice)) // for arrays, cap and len are equal

	// slices get there own indeces when created from array;
	// if underlying slices element is updated, it also gets updated in the array
	gradeSlice := grades[:3] // contains elements with index 0 to 2 from grade array
	fmt.Println(grades)
	fmt.Println(gradeSlice)

	gradeSlice[1] = 13
	fmt.Println("after modification")
	fmt.Println(grades)
	fmt.Println(gradeSlice)
	fmt.Println("")

	// append elements to a slice using append function
	anArray := [4]int{1, 2, 3, 4}
	aSlice := anArray[1:3]
	fmt.Println(aSlice)
	fmt.Println(len(aSlice))
	fmt.Println(cap(aSlice))
	fmt.Println("")

	// if not enough capacity in slice, append function will double capacity and kind of create a new array
	aSlice = append(aSlice, 900, -90, 50)
	fmt.Println(aSlice)
	fmt.Println(len(aSlice))
	fmt.Println(cap(aSlice))
	fmt.Println("")

	// append slice to slice
	arr1 := [4]int{0, 1, 2, 3}
	arrSlice1 := arr1[:2]
	arr2 := [4]int{4, 5, 6, 7}
	arrSlice2 := arr2[2:]
	newSlice = append(arrSlice1, arrSlice2...)
	fmt.Println(newSlice)
	fmt.Println("")

	// delete elements from slice
	arr1 = [4]int{0, 1, 2, 3}
	elementIndex := 2
	arrSlice1 = arr1[:elementIndex]
	arrSlice2 = arr1[elementIndex+1:]
	newSlice = append(arrSlice1, arrSlice2...)
	fmt.Println(newSlice)

	// copy from a slice with copy function (must be same datatype)
	srcSlice := []int{1, 2, 3}
	destSlice := make([]int, 3)
	numbers := copy(destSlice, srcSlice)
	fmt.Println(destSlice)
	fmt.Println("copied elements: ", numbers)

	// slices and loops
	loopArr := []int{10, 20, 30}

	for index, value := range loopArr {
		fmt.Println(index, "=>", value)
	}

	// if index is not needed also possible this way
	for _, value := range loopArr {
		fmt.Println(value)
	}

}
