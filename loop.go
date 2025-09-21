package main

import "fmt"

func loop() {
	// 1) Simple for loop (counter-based)
	// 👉 Runs a fixed number of times (1 to 5).
	fmt.Println("1) Simple for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i) // prints numbers 1 to 5
	}

	// 2) While-style loop (condition-based)
	// 👉 Go does not have a separate while loop,
	// but we can use for + condition like a while loop.
	fmt.Println("\n2) While-style loop:")
	x := 1
	for x <= 5 {
		fmt.Println(x) // prints numbers 1 to 5
		x++
	}

	// 3) Infinite loop (stopped with break)
	// 👉 "for {}" is an infinite loop. We stop it using break.
	fmt.Println("\n3) Infinite loop with break:")
	y := 1
	for {
		fmt.Println("Value of y:", y)
		y++
		if y > 3 {
			break // loop stops when y > 3
		}
	}

	// 4) Range loop on slice
	// 👉 Used to iterate over all elements in a slice/array.
	fmt.Println("\n4) Range loop on slice:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 5) Range loop on map
	// 👉 Used to iterate over all key-value pairs in a map.
	fmt.Println("\n5) Range loop on map:")
	students := map[string]int{"Amit": 18, "Kunal": 20, "Ravi": 22}
	for name, age := range students {
		fmt.Println("Name:", name, "Age:", age)
	}
}
