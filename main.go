package main

import (
	"errors"
	"fmt"
	"strconv"
)

func findIndex(s []int, num int) (int, error) {
	for i, n := range s {
		if n == num {
			return i, nil
		}
	}
	errorMsg := "Number not found index " + strconv.Itoa(num)
	return 0, errors.New(errorMsg)
}

func main() {
	i, err := findIndex([]int{1, 2, 3}, 20)
	if err != nil {
		fmt.Println("err :", err)
	} else {
		fmt.Println("index :", i)
	}
}
