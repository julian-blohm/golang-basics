package main

import "fmt"

// beschrieben am beispiel berechnen der fakultät einer zahl bspw 5! ist 5*4*3*2*1
func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// another example; prints numbers from a number until 1
func printNumbersUntil1(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n) // if switching this line an the next the function would print from 1 to 5
	print(n - 1)
}

func main() {
	n := 5
	result := factorial(n)

	fmt.Println("Factorial of", n, "is:", result)

	print(5)

}
