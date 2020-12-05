package main

import "fmt"

func main() {
	courses := []string{
		"Java",
		"C++",
		"Python",
	}

	for k, v := range courses {
		fmt.Println("k", k, " v : ", v)
	}

	fmt.Println("-----------------")
	credits := map[string]int{
		"Java":   3,
		"C++":    3,
		"Python": 3,
	}

	for k, v := range credits {
		fmt.Println("k", k, " v : ", v)
	}

}
