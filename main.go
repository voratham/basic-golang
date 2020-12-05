package main

import "fmt"

func add(a, b int) int {
	return a + b

}

func next(num int) (int, int) {
	return num, num + 1
}

func main() {
	result := add(1, 1)
	current, next := next(1)
	fmt.Println("result ::", result)
	fmt.Println(current, next)
}
