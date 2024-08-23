package main

import "fmt"

func main() {
	/*
	* Comparison operators
	* equal: ==
	* not equal: !=
	* less than: <
	* greater than: >
	* less or equal: <=
	* greater or equal: >=
	 */
	var name1 string = "bob"
	var name2 string = "joe"
	fmt.Println(name1 == name2)
	fmt.Println(name1 != name2)

	var number1, number2 = 2, 3
	fmt.Println(number1 < number2)
	fmt.Println(number1 <= number2)
	fmt.Println(number1 > number2)
	fmt.Println(number1 >= number2)

	/*
	* Arithmetic operators
	* addition: + ; with numbers an strings possible
	* subtraction: -
	* multiplication: *
	* division: /
	* modulus: % ; returns remainder when left operand is divided by right operand
	* increment: ++ ; unary operator, does not need two operands; increments operand by one
	* decrement: -- ; unary operator, does not need two operands
	 */

	fmt.Println(name1 + name2)
	fmt.Println(number1 - number2)
	fmt.Println(number1 * number2)
	fmt.Println(number1 / number2)
	fmt.Println(number1 % number2)

	number1++
	fmt.Println(number1)
	number2--
	fmt.Println(number1)

}
