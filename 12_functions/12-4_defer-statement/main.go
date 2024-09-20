package main

import "fmt"

/* defer statement delays execution of a function until surrounding function returns
* deferred call's arguments are evaluated immediately, function not executed yet
 */

func printName(str string) {
	fmt.Println(str)
}

func printAge(age int) {
	fmt.Println(age)
}

func printAdress(adr string) {
	fmt.Println(adr)
}

func main() {
	printName("Joe")
	defer printAge(23)
	printAdress("my-street 187")
	defer printAge(24)
}
