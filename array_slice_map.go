package main

import "fmt"

func array() {
	// ------------------------------
	// 1. Array Example
	// ------------------------------
	// Array = fixed size data structure
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println("Array elements:", arr)

	// ------------------------------
	// 2. Slice Example
	// ------------------------------
	// Slice = dynamic, flexible view of an array
	slice := arr[1:4] // from index 1 to 3 (4 is excluded)
	fmt.Println("Slice from array:", slice)

	// We can append new elements to a slice
	slice = append(slice, 60, 70)
	fmt.Println("Slice after append:", slice)

	// ------------------------------
	// 3. Map Example
	// ------------------------------
	// Map = key-value pairs
	studentMarks := make(map[string]int)

	// Insert values
	studentMarks["Kunal"] = 85
	studentMarks["Amit"] = 90
	studentMarks["Sneha"] = 78

	fmt.Println("Student Marks:", studentMarks)

	// Access a value by key
	fmt.Println("Kunal's Marks:", studentMarks["Kunal"])

	// Check if a key exists
	mark, exists := studentMarks["Rohit"]
	if exists {
		fmt.Println("Rohit's Marks:", mark)
	} else {
		fmt.Println("Rohit's Marks not found")
	}
}
