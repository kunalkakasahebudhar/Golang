package main

import "fmt"

func input_output() {
	// ------------------------------
	// Output example
	// ------------------------------
	fmt.Println("Hello, welcome to Go Input/Output example!")
	fmt.Print("This is printed using Print (no newline). ")
	fmt.Println("This is printed using Println (with newline).")
	fmt.Printf("Formatted number: %d, string: %s\n", 42, "GoLang") // printf style

	// ------------------------------
	// Input example
	// ------------------------------
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // take string input

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age) // take integer input

	fmt.Println("Hello,", name, "you are", age, "years old.")

	// ------------------------------
	// Input multiple values
	// ------------------------------
	var x, y int
	fmt.Print("Enter two numbers separated by space: ")
	fmt.Scan(&x, &y) // scan multiple inputs
	fmt.Println("Sum:", x+y)
}
