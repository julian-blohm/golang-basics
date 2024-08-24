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
	fmt.Print("equal: ")
	fmt.Println(name1 == name2)
	fmt.Print("not equal: ")
	fmt.Println(name1 != name2)

	var number1, number2 = 2, 3
	fmt.Print("less than: ")
	fmt.Println(number1 < number2)
	fmt.Print("less or equal: ")
	fmt.Println(number1 <= number2)
	fmt.Print("greater than: ")
	fmt.Println(number1 > number2)
	fmt.Print("greater or equal: ")
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

	fmt.Print("addition: ")
	fmt.Println(name1 + name2)
	fmt.Print("subtraction: ")
	fmt.Println(number1 - number2)
	fmt.Print("multiplication: ")
	fmt.Println(number1 * number2)
	fmt.Print("division: ")
	fmt.Println(number1 / number2)
	fmt.Print("modulus: ")
	fmt.Println(number1 % number2)

	number1++
	fmt.Print("increment: ")
	fmt.Println(number1)
	number1--
	fmt.Print("decrement: ")
	fmt.Println(number1)

	/*
	* Logical operator
	* logical and: &&
	* logical or: ||
	* logical not: ! ; unary operator
	 */

	number1 = 10
	fmt.Println("logical AND: ")
	fmt.Println((number1 < 100) && (number1 < 200))
	fmt.Println((number1 < 300) && (number1 < 0))

	fmt.Println("logical OR: ")
	fmt.Println((number1 < 100) || (number1 < 0))
	fmt.Println((number1 > 100) || (number1 < 0))

	// just for showcasing; sonar doesnt like it
	fmt.Print("logical NOT: ")
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
	fmt.Println("assign: ", number1)

	number1 = 5
	number1 += number2 // same as num1 = num1 + num2
	fmt.Println("add and assign: ", number1)

	number1 = 5
	number1 -= number2 // same as num1 = num1 - num2
	fmt.Println("subtract and assign: ", number1)

	number1 = 5
	number1 *= number2 // same as num1 = num1 * num2
	fmt.Println("multiply and assign: ", number1)

	number1 = 5
	number1 /= number2 // same as num1 = num1 / num2
	fmt.Println("divide and assign quotient: ", number1)

	number1 = 5
	number1 %= number2 // same as num1 = num1 % num2 (remainder of division is asigned)
	fmt.Println("divide and assign modulus: ", number1)

	/*
	* Bitwise operators
	* bitwise and: & ; takes two numbers as operands and does AND on every bit of two numbers
	* bitwise or: | ; similar bitwise AND but with OR
	* bitwise xor: ^ ; XOR on every bit of the number
	* left shift: << ; shift all bits left by specified number
	* right shift: >>
	 */
	// bitwise AND
	var x, y int = 12, 25 // 00001100 and 00011001 in binary
	z := x & y            // 8  (00001000) would be the anser
	fmt.Println("bitwise AND: ", z)

	// bitwise OR
	z = x | y // 29 (00011101)
	fmt.Println("bitwise OR: ", z)

	// bitwise XOR
	z = x ^ y // 21 (00010101)
	fmt.Println("bitwise XOR: ", z)

	// shift left
	z = x << 1 // 24 (00011000)
	fmt.Println("shift left: ", z)

	// shift right
	z = x >> 2 // 3 (00000011)
	fmt.Println("shift right: ", z)
}
