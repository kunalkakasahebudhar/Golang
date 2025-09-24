package main

import "fmt"

func Operators() {
	a := 10
	b := 20

	// Arithmetic Operators: + add, - subtract, * multiply, / divide, % modulus
	fmt.Println("a + b =", a+b) // Addition
	fmt.Println("a - b =", a-b) // Subtraction
	fmt.Println("a * b =", a*b) // Multiplication
	fmt.Println("a / b =", a/b) // Division
	fmt.Println("a % b =", a%b) // Modulus (remainder)

	// Relational Operators: == equal, != not equal, > greater, < less, >= greater/equal, <= less/equal
	fmt.Println("a == b:", a == b) // Equal
	fmt.Println("a != b:", a != b) // Not equal
	fmt.Println("a > b:", a > b)   // Greater than
	fmt.Println("a < b:", a < b)   // Less than
	fmt.Println("a >= b:", a >= b) // Greater or equal
	fmt.Println("a <= b:", a <= b) // Less or equal

	// Logical Operators: && and, || or, ! not
	x, y := true, false
	fmt.Println("x && y:", x && y) // AND
	fmt.Println("x || y:", x || y) // OR
	fmt.Println("!x:", !x)         // NOT

	// Assignment Operators: = assign, += add/assign, -= subtract/assign, *= multiply/assign, /= divide/assign, %= modulus/assign
	c := 5
	c += 2 // c = c + 2
	fmt.Println("c after += 2:", c)
	c -= 1 // c = c - 1
	fmt.Println("c after -= 1:", c)
	c *= 3 // c = c * 3
	fmt.Println("c after *= 3:", c)
	c /= 2 // c = c / 2
	fmt.Println("c after /= 2:", c)
	c %= 2 // c = c % 2
	fmt.Println("c after %= 2:", c)

	// Unary Operators: ++ increment, -- decrement, + positive, - negative, ! not
	d := 5
	d++ // increment by 1
	fmt.Println("d after d++:", d)
	d-- // decrement by 1
	fmt.Println("d after d--:", d)
	fmt.Println("Unary -d:", -d) // negative

	// Bitwise Operators: & AND, | OR, ^ XOR, &^ AND NOT, << left shift, >> right shift
	e, f := 5, 3                 // 5=101, 3=011 in binary
	fmt.Println("e & f:", e&f)   // AND
	fmt.Println("e | f:", e|f)   // OR
	fmt.Println("e ^ f:", e^f)   // XOR
	fmt.Println("e &^ f:", e&^f) // AND NOT
	fmt.Println("e << 1:", e<<1) // Left shift
	fmt.Println("e >> 1:", e>>1) // Right shift
}

// Explanation

// Arithmetic Operators (+ - * / %) → Used to perform basic math on numbers.

// Relational Operators (== != > < >= <=) → Used to compare two values, result is true or false.

// Logical Operators (&& || !) → Used to combine or negate boolean values.

// Assignment Operators (= += -= *= /= %=) → Used to assign values to variables, sometimes with calculation.

// Unary Operators (++ -- + - !) → Used with a single operand to increment, decrement, or negate.

// Bitwise Operators (& | ^ &^ << >>) → Used to perform operations on the binary form of numbers.
