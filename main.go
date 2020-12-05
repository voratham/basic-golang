package main

import "fmt"

func main() {
	// array type
	arr := [5]string{"A", "B", "C", "D", "E"}

	// slice type
	result := arr[2:4]

	// arr[0] = "X"
	result[0] = "M" // ☝️ in case  slice not equal  arr

	fmt.Println("arr ::", arr)
	fmt.Println("result ::", result)

}
