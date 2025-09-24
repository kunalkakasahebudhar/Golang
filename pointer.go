package main

import "fmt"

// Function to swap two integers using pointers
func swap(x *int, y *int) {
	temp := *x // store value at x
	*x = *y    // assign value at y to x
	*y = temp  // assign temp to y
}

func pointer() {
	a := 10
	b := 20

	fmt.Println("Before swapping: a =", a, ", b =", b)

	// Pass addresses of a and b to the function
	swap(&a, &b)

	fmt.Println("After swapping: a =", a, ", b =", b)
}
