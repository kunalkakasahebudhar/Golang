package main

import (
	"fmt"
	"strings"
)

func stringfun() {
	// ------------------------------
	// One-line definitions
	// ------------------------------
	fmt.Println("String: A sequence of characters.") // One-line definition

	// ------------------------------
	// 1. String Declaration
	// ------------------------------
	var str1 string = "Hello"
	str2 := "World"

	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)

	// ------------------------------
	// 2. String Concatenation
	// ------------------------------
	combined := str1 + " " + str2
	fmt.Println("Concatenated String:", combined)

	// ------------------------------
	// 3. String Length
	// ------------------------------
	fmt.Println("Length of combined string:", len(combined))

	// ------------------------------
	// 4. Convert to Uppercase / Lowercase
	// ------------------------------
	fmt.Println("Uppercase:", strings.ToUpper(combined))
	fmt.Println("Lowercase:", strings.ToLower(combined))

	// ------------------------------
	// 5. Check if string contains a substring
	// ------------------------------
	fmt.Println("Contains 'Hello'?", strings.Contains(combined, "Hello"))
	fmt.Println("Contains 'Go'?", strings.Contains(combined, "Go"))

	// ------------------------------
	// 6. Find index of a substring
	// ------------------------------
	fmt.Println("Index of 'World':", strings.Index(combined, "World"))

	// ------------------------------
	// 7. Replace substring
	// ------------------------------
	newStr := strings.Replace(combined, "World", "Golang", 1)
	fmt.Println("After Replace:", newStr)

	// ------------------------------
	// 8. Split string into slice
	// ------------------------------
	words := strings.Split(combined, " ")
	fmt.Println("Split words:", words)

	// ------------------------------
	// 9. Join slice into string
	// ------------------------------
	joined := strings.Join(words, "-")
	fmt.Println("Joined string with '-':", joined)

	// ------------------------------
	// 10. Trim spaces
	// ------------------------------
	str3 := "   GoLang   "
	fmt.Println("Before Trim:", str3)
	fmt.Println("After Trim:", strings.TrimSpace(str3))
}
