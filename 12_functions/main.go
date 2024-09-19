package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

// function syntax: func <function name>(<params>) <return type> { <body of function> }
func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

// possible to return mutliple values; can also be named directly
func operation(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// variadic function => variable number of arguments of same datatype
func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func printNameAndFriends(name string, friends ...string) {
	fmt.Println("hey", name, ", these are my friends: ")
	for _, friend := range friends {
		fmt.Printf("%s, ", friend)
	}
}

// using blank identifier
func functionThing() (int, int) {
	return 42, 50
}

func main() {

	sayHello("jon")

	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)

	sum, diff := operation(20, 10)
	fmt.Println(sum, diff)

	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10, 20, 30))

	printNameAndFriends("Joe", "John", "James")

	// using blank identifier if there are multiple return values but we only want a specific
	variableThing, _ := functionThing()
	fmt.Println(variableThing)

}
