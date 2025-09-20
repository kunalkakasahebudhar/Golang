package main

import "fmt"

// Constants are already defined at package level
const appName = "GoLangApp"
const version float64 = 1.0
const maxUsers = 100

const (
	Pi  = 3.14159
	E   = 2.71828
	URL = "https://golang.org"
)

const (
	A = iota // 0
	B        // 1
	C        // 2
)

func Constants() {
	fmt.Println("App Name:", appName)
	fmt.Println("Version:", version)
	fmt.Println("Max Users Allowed:", maxUsers)

	fmt.Println("Pi:", Pi)
	fmt.Println("Euler's Number:", E)
	fmt.Println("Go Website:", URL)

	fmt.Println("Iota values:", A, B, C)
}
