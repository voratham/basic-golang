package main

import "fmt"

// func inc(num int) int {
// 	return num + 1
// }

func inc(num *int) {
	*num = *num + 1
}

func main() {
	i := 20
	fmt.Println("first i ", i)
	inc(&i)
	fmt.Println("seconds i ", i)

}
