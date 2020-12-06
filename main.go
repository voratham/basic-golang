package main

import "fmt"

func sum(ch chan int, nums []int) {
	result := 0
	for _, n := range nums {
		result += n
	}
	ch <- result
}

func main() {
	ch := make(chan int)
	result := 0
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}

	go sum(ch, nums[0:10])
	go sum(ch, nums[10:20])
	go sum(ch, nums[20:])

	for count := 0; count < 3; count++ {
		result += <-ch
	}
	fmt.Println("final result :", result)
}
