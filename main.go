package main

import "fmt"

func main() {
	day := "Saturday"

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend 😎")

	default:
		fmt.Println("Workday 😭😭😭 !!")
	}
}
