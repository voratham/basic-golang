package main

import "fmt"

func main() {
	words := []string{"Hello", "Hi", "Bye", "Thailand", "Japan"}
	fmt.Println("1 words ::", words)

	// Append new value to slice
	words = append(words, "Go")
	fmt.Println("2 append words ::", words)

	// Remove value  on index => 2(Bye)
	// {"Hello", "Hi"} + {"Thailand", "Japan"}

	// words = append(words[:2], "Thailand", "Japan")
	words = append(words[:2], words[3:]...) // spread operator
	fmt.Println("3 remove words ::", words)
}
