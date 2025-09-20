package main

import "fmt"

func datatype() {
	// Boolean is used to store true or false values
	var isAvailable bool = true
	var isDone bool
	fmt.Println("Boolean:", isAvailable, isDone)

	// Integer is used to store whole numbers
	var age int = 18
	var score uint8 = 255
	fmt.Println("Integers:", age, score)

	// Float is used to store decimal numbers
	var price float32 = 99.99
	var pi float64 = 3.14159
	fmt.Println("Floats:", price, pi)

	// Complex is used to store complex numbers with real and imaginary parts
	var c complex64 = 2 + 3i
	fmt.Println("Complex:", c)

	// String is used to store text
	var name string = "Kunal"
	var empty string
	fmt.Println("Strings:", name, empty)

	// Byte is used for small numbers or raw data, Rune is used for Unicode characters
	var b byte = 'A'
	var r rune = 'अ'
	fmt.Println("Byte:", b, "Rune:", r)

	// Array is used to store fixed-size collection of same type values
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	// Slice is used to store dynamic-size collection of values
	nums := []int{10, 20, 30}
	nums = append(nums, 40)
	fmt.Println("Slice:", nums)

	// Map is used to store key-value pairs
	ages := map[string]int{
		"Kunal": 18,
		"Rohit": 20,
	}
	ages["Maya"] = 22
	fmt.Println("Map:", ages)

	// Struct is used to group different types of data together
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Kunal", Age: 18}
	fmt.Println("Struct:", p)

	// Pointer is used to store address of a variable
	x := 100
	px := &x
	fmt.Println("Pointer:", *px)

	// Function type is used to store a function as a variable
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println("Function:", add(5, 10))

	// Interface is used to define behavior that other types can implement
	var s Shape = Circle{Radius: 5}
	fmt.Println("Interface (Circle Area):", s.Area())

	// Channel is used for communication between goroutines
	ch := make(chan int, 1)
	ch <- 42
	fmt.Println("Channel:", <-ch)

	// Nil is used to represent zero value for pointers, maps, slices, interfaces, channels
	var ptr *int = nil
	var m map[string]int = nil
	fmt.Println("Nil pointer and map:", ptr, m)
}

// Interface is used to define behavior that other types can implement
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}
