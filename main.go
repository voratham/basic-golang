package main

import "fmt"

func main() {
	credits := map[string]int{
		"Java":   3,
		"C++":    3,
		"Python": 4,
	}

	delete(credits, "Python")
	credits["C#"] = 3
	result, ok := credits["C#"]
	fmt.Println(result, ok)

	// credits := make(map[string]int)
	// credits["Java"] = 3
	// credits["C++"] = 3
	// credits["Python"] = 3

	// result, ok := credits["Python"]
	// fmt.Println(result, ok)

	// fmt.Print(credits)

}
