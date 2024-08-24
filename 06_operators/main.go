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
	number1--
	fmt.Println(number1)

	/*
	* Logical operator
	* logical and: &&
	* logical or: ||
	* logical not: ! ; unary operator
	 */

	number1 = 10
	fmt.Println((number1 < 100) && (number1 < 200))
	fmt.Println((number1 < 300) && (number1 < 0))
	fmt.Println("")

	fmt.Println((number1 < 100) || (number1 < 0))
	fmt.Println((number1 > 100) || (number1 < 0))

	// just for showcasing; sonar doesnt like it
	fmt.Println(!(true))

	/*
	* Assignment operators
	* assign: =
	* add and assign: +=
	* subtract and assign: -=
	* multiply and assign: *=
	* divide and assign quotient: /=
	* divide and assign modulus: %=
	 */
	number1 = 5
	number2 = 10
	number1 = number2
	fmt.Println(number1)

	number1 = 5
	number1 += number2 // same as num1 = num1 + num2
	fmt.Println(number1)

	number1 = 5
	number1 -= number2 // same as num1 = num1 - num2
	fmt.Println(number1)

	number1 = 5
	number1 *= number2 // same as num1 = num1 * num2
	fmt.Println(number1)

	number1 = 5
	number1 /= number2 // same as num1 = num1 / num2
	fmt.Println(number1)

	number1 = 5
	number1 %= number2 // same as num1 = num1 % num2 (remainder of division is asigned)
	fmt.Println(number1)

}
