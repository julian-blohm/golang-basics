package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("How old are you?")
	fmt.Scanf("%d", &age)

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age < 18 {
		fmt.Println("Your are a minor")
	} else {
		fmt.Println("no valid number")
	}

	var mybool bool = true
	switch mybool {
	case true:
		fmt.Println("this is true")
	case false:
		fmt.Println("this is false")
	default:
		fmt.Println("this is not a boolean")
	}

	switch {
	case age <= 0:
		fmt.Println("Enter a valid number")
	case age < 18:
		fmt.Println("You are a minor")
	case age >= 18:
		fmt.Println("Your are a minor")
	default:
		fmt.Println("Did you enter a number?")
	}
}
